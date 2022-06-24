package workspace

import (
	"io/ioutil"

	"github.com/playmean/guest/settings"
)

func (w *Workspace) Load() error {
	f, err := w.VirtualFs.Open(SettingsPath)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	err = w.Folder.Load(data, settings.FormatJson)
	if err != nil {
		return err
	}

	return nil
}
