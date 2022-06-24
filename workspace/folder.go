package workspace

import (
	"github.com/playmean/guest/settings"
)

type Folder struct {
	Description string `json:"description"`

	Variables   map[string]string      `json:"variables,omitempty" yaml:"variables,omitempty"`
	Scripts     map[string]string      `json:"scripts,omitempty" yaml:"scripts,omitempty"`
	HandOptions map[string]interface{} `json:"options,omitempty" yaml:"options,omitempty"`
}

func (f *Folder) Load(data []byte, format settings.SettingsFormat) error {
	err := settings.Parse(data, f, format)
	if err != nil {
		return err
	}

	return nil
}

func (w *Workspace) LoadFolders(endPath string) ([]*Folder, error) {
	folders := make([]*Folder, 0)

	err := w.loadFoldersRecursively(endPath, &folders)
	if err != nil {
		return nil, err
	}

	return folders, nil
}
