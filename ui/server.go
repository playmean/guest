package ui

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/playmean/guest/workspace"
)

//go:embed frontend/dist/*
var embedded embed.FS

type Server struct {
	app          *fiber.App
	workspace    *workspace.Workspace
	externalVars map[string]string
}

// TODO 24.06.2022 extend from error
type ServerError struct {
	Error string `json:"error"`
}

func NewServer(w *workspace.Workspace, externalVars map[string]string) *Server {
	app := fiber.New()
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(embedded),
		PathPrefix: "frontend/dist",
	}))

	s := new(Server)
	s.app = app
	s.workspace = w
	s.externalVars = externalVars

	api := app.Group("/api")
	api.Get("/version", s.methodGetVersion)
	api.Get("/workspace", s.methodGetWorkspace)

	return s
}

func (s *Server) Start(address string) error {
	return s.app.Listen(address)
}
