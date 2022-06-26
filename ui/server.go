package ui

import (
	"embed"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
	osSignal     chan os.Signal
	doneSignal   chan error
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
	s.osSignal = make(chan os.Signal)
	s.doneSignal = make(chan error)

	api := app.Group("/api")
	api.Get("/version", s.methodGetVersion)
	api.Get("/workspace", s.methodGetWorkspace)

	signal.Notify(s.osSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	return s
}

func (s *Server) Start(address string) error {
	go func() {
		err := s.app.Listen(address)

		s.doneSignal <- err
	}()

	select {
	case <-s.osSignal:
		return nil
	case err := <-s.doneSignal:
		return err
	}
}
