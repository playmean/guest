package workspace

import "guest/settings"

func (w *Workspace) Save() error {
	err := w.VirtualFs.MkdirAll(".guest", 0655)
	if err != nil {
		return err
	}

	err = w.writeSettings()
	if err != nil {
		return err
	}

	err = w.writeMisc()
	if err != nil {
		return err
	}

	return nil
}

func (w *Workspace) writeSettings() error {
	settingsFile, err := w.VirtualFs.Create(SettingsPath)
	if err != nil {
		return err
	}

	defer settingsFile.Close()

	data, err := settings.Stringify(w.Folder, settings.FormatJson)
	if err != nil {
		return err
	}

	_, err = settingsFile.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (w *Workspace) writeMisc() error {
	ignoreFile, err := w.VirtualFs.Create(".guest/.gitignore")
	if err != nil {
		return err
	}

	defer ignoreFile.Close()

	// TODO use .local directories
	_, err = ignoreFile.Write([]byte(".local/"))
	if err != nil {
		return err
	}

	return nil
}
