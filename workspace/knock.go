package workspace

import (
	"guest/knock"
	"path/filepath"

	"github.com/imdario/mergo"
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

	externalScripts := make(map[string]string)
	handOptions := make(map[string]interface{})

	for _, folder := range folders {
		for scriptType, scriptPath := range folder.Scripts {
			externalScripts[scriptType] = scriptPath
		}

		err = mergo.Merge(&handOptions, folder.HandOptions, mergo.WithOverride)
		if err != nil {
			return nil, err
		}
	}

	err = k.RunScript(externalScripts, vars, "before")
	if err != nil {
		return nil, err
	}

	err = k.PatchOptions(handOptions[k.GetType()])
	if err != nil {
		return nil, err
	}

	err = k.ApplyVariables(vars)
	if err != nil {
		return nil, err
	}

	return k.Run()
}
