package ui

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/playmean/guest/workspace"
)

type GetWorkspaceResponse struct {
	Description string               `json:"description"`
	Variables   map[string]string    `json:"variables"`
	Tree        []WorkspaceTreeEntry `json:"tree"`
}

type OpenWorkspaceRequest struct {
	Path string `json:"path"`
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
		return c.Status(http.StatusNotFound).JSON(ServerError{
			Error: "no workspace loaded",
		})
	}

	return s.sendWorkspaceInfo(c)
}

func (s *Server) methodOpenWorkspace(c *fiber.Ctx) error {
	var err error

	req := new(OpenWorkspaceRequest)

	err = c.BodyParser(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(ServerError{
			Error: "cannot parse request",
		})
	}

	if req.Path == "" {
		return c.Status(http.StatusBadRequest).JSON(ServerError{
			Error: "path must not be empty",
		})
	}

	s.workspace, err = workspace.FromStorage("git", req.Path)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(ServerError{
			Error: "cannot load workspace",
		})
	}

	return s.sendWorkspaceInfo(c)
}

func (s *Server) methodCreateWorkspace(c *fiber.Ctx) error {
	var err error

	s.workspace, err = workspace.NewWorkspace()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(ServerError{
			Error: "cannot create workspace",
		})
	}

	return s.sendWorkspaceInfo(c)
}
