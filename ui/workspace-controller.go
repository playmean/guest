package ui

import (
	"net/http"
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/playmean/guest/storage"
)

type GetWorkspaceResponse struct {
	Description string               `json:"description"`
	Variables   map[string]string    `json:"variables"`
	Tree        []WorkspaceTreeEntry `json:"tree"`
}

type WorkspaceTreeEntryType string

const (
	WorkspaceTreeEntryTypeDir  WorkspaceTreeEntryType = "dir"
	WorkspaceTreeEntryTypeFile WorkspaceTreeEntryType = "file"
)

type WorkspaceTreeEntry struct {
	Path     string               `json:"path"`
	Title    string               `json:"title"`
	Type     string               `json:"type"`
	Children []WorkspaceTreeEntry `json:"children,omitempty"`
}

func (s *Server) methodGetWorkspace(c *fiber.Ctx) error {
	if s.workspace == nil {
		c.JSON(ServerError{
			Error: "no workspace",
		})

		return c.SendStatus(404)
	}

	paths, err := storage.ScanDirectory(s.workspace.PathInfo.Dir, s.workspace.VirtualFs, &storage.ScanOptions{
		Ignore:             []string{".git"},
		IncludeDirectories: true,
		MarkDirectories:    true,
	})
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(GetWorkspaceResponse{
		Description: s.workspace.Description,
		Variables:   s.workspace.Variables,
		Tree:        buildWorkspaceTree("", paths),
	})
}

func buildWorkspaceTree(parentPath string, paths []string) []WorkspaceTreeEntry {
	chidlrenEntries := make([]WorkspaceTreeEntry, 0)
	lastDir := ""

	for _, entryPath := range paths {
		if entryPath == parentPath {
			continue
		}

		if !strings.HasPrefix(entryPath, parentPath) {
			continue
		}

		if lastDir != "" && strings.HasPrefix(entryPath, lastDir) {
			continue
		}

		if strings.HasSuffix(entryPath, "/") {
			chidlrenEntries = append(chidlrenEntries, WorkspaceTreeEntry{
				Path:     entryPath,
				Title:    path.Base(entryPath),
				Type:     "dir",
				Children: buildWorkspaceTree(entryPath, paths),
			})

			lastDir = entryPath

			continue
		}

		chidlrenEntries = append(chidlrenEntries, WorkspaceTreeEntry{
			Path:  entryPath,
			Title: path.Base(entryPath),
			Type:  "file",
		})
	}

	return chidlrenEntries
}
