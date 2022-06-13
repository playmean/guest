package workspace

import (
	"guest/settings"
	"guest/storage"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func (w *Workspace) loadFoldersRecursively(endPath string, folders *[]*Folder) error {
	folder, _ := w.loadFolderForPath(endPath)

	*folders = append(*folders, folder)

	prevDir := strings.TrimSuffix(endPath, "/"+filepath.Base(endPath))

	if prevDir == endPath {
		return nil
	}

	return w.loadFoldersRecursively(prevDir, folders)
}

func (w *Workspace) loadFolderForPath(endPath string) (*Folder, error) {
	folderPath, err := storage.ScanParent(endPath, w.VirtualFs, ".guest.json")
	if err != nil {
		return nil, err
	}

	localFolderPath, _ := storage.ScanParent(endPath, w.VirtualFs, ".local.json")
	if localFolderPath != "" {
		folderPath = localFolderPath
	}

	f, err := w.VirtualFs.Open(folderPath)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	folder := new(Folder)

	err = folder.Load(data, settings.FormatJson)
	if err != nil {
		return nil, err
	}

	return folder, nil
}
