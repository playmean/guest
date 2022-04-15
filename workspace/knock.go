package workspace

import (
	"guest/knock"
	"path/filepath"
)

func (w *Workspace) Knock(path string, externalVars map[string]string) (*knock.Result, error) {
	k, err := knock.FromFile(path, w.VirtualFs)
	if err != nil {
		return nil, err
	}

	knockDir := filepath.Dir(path)

	folders, err := w.LoadFolders(knockDir)
	if err != nil {
		return nil, err
	}

	vars, err := w.LoadVariables(folders, externalVars)
	if err != nil {
		return nil, err
	}

	err = k.RunScript(vars, "before")
	if err != nil {
		return nil, err
	}

	err = k.ApplyVariables(vars)
	if err != nil {
		return nil, err
	}

	return k.Run()
}
