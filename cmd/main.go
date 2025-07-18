//go:build wireinject
// +build wireinject

package main

import (
	"log"
	"wowzaman106/config"
	"wowzaman106/infrastructure/server"
	"wowzaman106/internal/handler/gamehdl"

	"github.com/gofiber/fiber/v3"
	"github.com/google/wire"
)

type dependencies struct {
	server *fiber.App
}

func initialize() *dependencies {
	wire.Build(
		server.New,
		gamehdl.New,
		wire.Struct(new(dependencies), "server"),
	)
	return nil
}

func main() {
	d := initialize()

	log.Printf("Server is running on port %s", config.Get().Server.Port)

	// Graceful shutdown and other server logic can be added here
	defer func() {
		if err := d.server.Shutdown(); err != nil {
			log.Printf("Error during shutdown: %s", err.Error())
		}
		log.Println("Server gracefully stopped")
	}()

	// Start the server on the configured port
	if err := d.server.Listen(":" + config.Get().Server.Port); err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
