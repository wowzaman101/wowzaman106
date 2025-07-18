package server

import (
	"wowzaman106/internal/handler/gamehdl"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func New(gh gamehdl.Handler) *fiber.App {

	app := fiber.New()

	// Health check route
	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Game handler route
	app.Post("/game", gh.Play)

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	return app
}
