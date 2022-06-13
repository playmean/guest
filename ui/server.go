package ui

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed frontend/dist/*
var embedded embed.FS

type Server struct {
	app *fiber.App
}

func NewServer() *Server {
	app := fiber.New()
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(embedded),
		PathPrefix: "frontend/dist",
	}))

	s := new(Server)
	s.app = app

	return s
}

func (s *Server) Start(address string) error {
	return s.app.Listen(address)
}
