package ui

import (
	"net/http"
	"path"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/playmean/guest/storage"
)

func (s *Server) sendWorkspaceInfo(c *fiber.Ctx) error {
	paths, err := storage.ScanDirectory(s.workspace.PathInfo.Dir, s.workspace.VirtualFs, &storage.ScanOptions{
		Ignore:             []string{".git"},
		IncludeDirectories: true,
		MarkDirectories:    true,
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(ServerError{
			Error: "error while scanning workspace files tree",
		})
	}

	return c.Status(http.StatusOK).JSON(GetWorkspaceResponse{
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
