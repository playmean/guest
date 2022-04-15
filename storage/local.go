package storage

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
)

type LocalStorage struct {
	billy.Filesystem
}

func (s LocalStorage) Load(path string) (Storage, *StoragePathInfo, error) {
	workDir, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}

	workspaceDir, err := ScanParent(workDir, s, ".guest")
	if err != nil {
		return nil, nil, err
	}

	workspaceDir = filepath.Dir(workspaceDir)

	path = strings.TrimPrefix(workDir, workspaceDir)

	root, err := s.Chroot(workspaceDir)
	if err != nil {
		return nil, nil, err
	}

	return LocalStorage{
			root,
		}, &StoragePathInfo{
			Dir: path,
		}, nil
}

func (s LocalStorage) Save(path string) (Storage, error) {
	return s, nil
}

func NewLocal() Storage {
	return LocalStorage{
		osfs.New("/"),
	}
}
