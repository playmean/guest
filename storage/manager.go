package storage

import "fmt"

type Manager struct {
	storages map[string]Storage
}

var DefaultManager = NewManager()

func NewManager() *Manager {
	return &Manager{
		storages: make(map[string]Storage),
	}
}

func (m *Manager) RegisterStorage(name string, storage Storage) error {
	if _, ok := m.storages[name]; ok {
		return fmt.Errorf("storage provider '%s' already registered", name)
	}

	m.storages[name] = storage

	return nil
}

func (m *Manager) ResolveStorage(name string) (Storage, error) {
	if storage, ok := m.storages[name]; ok {
		return storage, nil
	}

	return nil, fmt.Errorf("storage provider '%s' not found", name)
}
