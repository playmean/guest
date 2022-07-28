package ui

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type GetVersionResponse struct {
	App string `json:"app"`
}

var VersionHash = "dev"

func (s *Server) methodGetVersion(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(GetVersionResponse{
		App: VersionHash,
	})
}
