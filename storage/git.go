package storage

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/go-git/go-git/v5/storage/memory"
	"golang.org/x/exp/slices"
)

type GitStorage struct {
	billy.Filesystem
}

func (s GitStorage) Load(path string) (Storage, *StoragePathInfo, error) {
	uri, err := url.Parse(path)
	if err != nil {
		return nil, nil, err
	}

	slashCount := strings.Count(uri.Path, "/")
	i := strings.LastIndexFunc(uri.Path, func(c rune) bool {
		if c == '/' {
			slashCount--

			return slashCount-1 == 0
		}

		return false
	})

	if i > 0 {
		uri.Path = uri.Path[:i]
	}

	if uri.Scheme == "" {
		uri.Scheme = "https"
	}

	if uri.Host == "" {
		uri.Host = "github.com"
	}

	repoUrl := uri.String()

	ref, err := getGitMainRef(repoUrl)
	if err != nil {
		return nil, nil, err
	}

	gitFs, _ := s.Chroot(".git")
	storage := filesystem.NewStorage(gitFs, cache.NewObjectLRUDefault())

	_, err = git.Clone(storage, s, &git.CloneOptions{
		URL:           repoUrl,
		SingleBranch:  true,
		ReferenceName: ref.Name(),
		Tags:          git.NoTags,
		// Progress:      os.Stdout,
	})
	if err != nil {
		return nil, nil, err
	}

	return s, &StoragePathInfo{
		Dir:        "",
		TrimPrefix: uri.Path,
	}, nil
}

func (s GitStorage) Save(path string) (Storage, error) {
	return s, nil
}

func NewGit() Storage {
	return GitStorage{
		memfs.New(),
	}
}

func getGitMainRef(repoPath string) (*plumbing.Reference, error) {
	mainRefNames := []string{"master", "main"}

	rem := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		Name: "origin",
		URLs: []string{repoPath},
	})

	refs, err := rem.List(&git.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, ref := range refs {
		refName := ref.Name().Short()

		if !ref.Name().IsBranch() {
			continue
		}

		if slices.Contains(mainRefNames, refName) {
			return ref, nil
		}
	}

	return nil, fmt.Errorf("reference not found (master/main)")
}
