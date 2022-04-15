package workspace

func (w *Workspace) LoadVariables(folders []*Folder, external map[string]string) (map[string]string, error) {
	vars := make(map[string]string, 0)

	for k, v := range external {
		vars[k] = v
	}

	for _, f := range folders {
		folderVars := f.Variables

		for k, v := range folderVars {
			if _, ok := vars[k]; !ok {
				vars[k] = v
			}
		}
	}

	for k, v := range w.Variables {
		if _, ok := vars[k]; !ok {
			vars[k] = v
		}
	}

	return vars, nil
}
