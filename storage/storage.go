package storage

import (
	"path/filepath"
	"strings"

	"github.com/go-git/go-billy/v5"
)

type StoragePathInfo struct {
	Dir        string
	TrimPrefix string
}

type Storage interface {
	billy.Filesystem

	Load(path string) (Storage, *StoragePathInfo, error)
	Save(path string) (Storage, error)
}

// TODO move to another file
func (i *StoragePathInfo) Resolve(path string) string {
	relativePath := strings.TrimPrefix(path, i.TrimPrefix)

	return filepath.Join(i.Dir, relativePath)
}
