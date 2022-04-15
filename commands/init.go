package commands

import (
	"fmt"
	"guest/storage"
	"guest/workspace"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func Init(c *cli.Context) error {
	exportStorage, _ := storage.DefaultManager.ResolveStorage("local")

	path := c.Args().First()

	exportPath, err := filepath.Abs(filepath.Base(path))
	if err != nil {
		return err
	}

	if storage.Exists(filepath.Join(exportPath, workspace.SettingsPath), exportStorage) {
		return fmt.Errorf("workspace is already exists")
	}

	w, err := workspace.NewWorkspace()
	if err != nil {
		return err
	}

	err = w.Export(exportPath, exportStorage, &workspace.ExportOptions{
		Ignore: []string{".git"},
	})
	if err != nil {
		return err
	}

	return nil
}
