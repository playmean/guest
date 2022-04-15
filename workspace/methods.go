package workspace

import (
	"fmt"
	"guest/storage"
)

type ExportOptions struct {
	Ignore []string
}

func (w *Workspace) Export(destPath string, destStorage storage.Storage, opts *ExportOptions) error {
	err := storage.CopyDirectory(".", destPath, w.VirtualFs, destStorage, &storage.CopyOptions{
		Ignore: opts.Ignore,
	})
	if err != nil {
		return err
	}

	w.VirtualFs, err = destStorage.Save(destPath)

	return err
}

func (w *Workspace) Validate() error {
	if !storage.Exists(SettingsPath, w.VirtualFs) {
		return fmt.Errorf("not a valid guest workspace")
	}

	return nil
}
