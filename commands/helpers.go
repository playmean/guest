package commands

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/playmean/guest/storage"
	"github.com/playmean/guest/workspace"

	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slices"
)

func resolveWorkspace(c *cli.Context, storageBlacklist []string) (*workspace.Workspace, error) {
	localStorage, _ := storage.DefaultManager.ResolveStorage("local")

	storageName := c.String("storage")
	loadPath := c.Args().First()

	if loadPath == "" {
		return nil, fmt.Errorf("download path must be provided")
	}

	if storageBlacklist != nil && slices.Contains(storageBlacklist, storageName) {
		return nil, fmt.Errorf("cannot use '%s' provider with this command", storageName)
	}

	if storageName == "" {
		loadFullPath, _ := filepath.Abs(loadPath)

		if storage.Exists(loadFullPath, localStorage) {
			storageName = "local"
		} else {
			storageName = "git"
		}
	}

	w, err := workspace.FromStorage(storageName, loadPath)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func resolveVariables(c *cli.Context) (map[string]string, error) {
	vars := make(map[string]string)

	varSlice := c.StringSlice("var")

	for _, varDef := range varSlice {
		k := ""
		v := ""

		unpackSlice(strings.Split(varDef, "="), &k, &v)

		if k == "" || v == "" {
			return nil, fmt.Errorf("invalid var definition: '%v'", varDef)
		}

		vars[k] = v
	}

	return vars, nil
}

func unpackSlice(s []string, vars ...*string) {
	for i, str := range s {
		if i >= len(vars) {
			return
		}

		*vars[i] = str
	}
}
