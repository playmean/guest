package storage

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
)

type MemoryStorage struct {
	billy.Filesystem
}

// func (s MemoryStorage) Open(filename string) (billy.File, error) {}

func (s MemoryStorage) Load(path string) (Storage, *StoragePathInfo, error) {
	return s, &StoragePathInfo{}, nil
}

func (s MemoryStorage) Save(path string) (Storage, error) {
	return s, nil
}

func NewMemory() Storage {
	return MemoryStorage{
		memfs.New(),
	}
}
