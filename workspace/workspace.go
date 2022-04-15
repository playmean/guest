package workspace

import (
	"guest/storage"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
)

type Workspace struct {
	Folder

	PathInfo  *storage.StoragePathInfo
	VirtualFs storage.Storage

	repository *git.Repository
	folders    []*Folder
}

const SettingsPath = ".guest/workspace.json"

func NewWorkspace() (*Workspace, error) {
	virtualFs := storage.NewMemory()
	gitFs, _ := virtualFs.Chroot(".git")
	gitStorage := filesystem.NewStorage(gitFs, cache.NewObjectLRUDefault())

	repo, err := git.Init(gitStorage, virtualFs)
	if err != nil {
		return nil, err
	}

	w := &Workspace{
		PathInfo:  &storage.StoragePathInfo{},
		VirtualFs: virtualFs,

		repository: repo,
		folders:    make([]*Folder, 0),
	}

	err = w.Save()
	if err != nil {
		return nil, err
	}

	return w, nil
}

// TODO add FromTemplate method

func FromStorage(storageName, path string) (*Workspace, error) {
	virtualFs, err := storage.DefaultManager.ResolveStorage(storageName)
	if err != nil {
		return nil, err
	}

	virtualFs, pathInfo, err := virtualFs.Load(path)
	if err != nil {
		return nil, err
	}

	gitFs, _ := virtualFs.Chroot(".git")
	gitStorage := filesystem.NewStorage(gitFs, cache.NewObjectLRUDefault())

	repo, err := git.Open(gitStorage, virtualFs)
	if err != nil {
		return nil, err
	}

	w := &Workspace{
		PathInfo:  pathInfo,
		VirtualFs: virtualFs,

		repository: repo,
		folders:    make([]*Folder, 0),
	}

	err = w.Validate()
	if err != nil {
		return nil, err
	}

	err = w.Load()
	if err != nil {
		return nil, err
	}

	// err = w.Scan(".")
	// if err != nil {
	// 	return nil, err
	// }

	return w, nil
}
