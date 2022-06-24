package ui

import (
	"github.com/gofiber/fiber/v2"
)

type GetVersionResponse struct {
	App string `json:"app"`
}

var VersionHash = "dev"

func (s *Server) methodGetVersion(c *fiber.Ctx) error {
	return c.JSON(GetVersionResponse{
		App: VersionHash,
	})
}
