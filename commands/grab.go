package commands

import (
	"fmt"
	"guest/storage"
	"guest/workspace"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func Grab(c *cli.Context) error {
	exportStorage, _ := storage.DefaultManager.ResolveStorage("local")

	w, err := resolveWorkspace(c, []string{"local"})
	if err != nil {
		return err
	}

	path := c.Args().First()

	if c.Args().Len() == 2 {
		path = c.Args().Get(1)
	}

	exportPath, err := filepath.Abs(filepath.Base(path))
	if err != nil {
		return err
	}

	if storage.Exists(exportPath, exportStorage) {
		return fmt.Errorf("directory '%s' already exists", exportPath)
	}

	return w.Export(exportPath, exportStorage, &workspace.ExportOptions{})
}
