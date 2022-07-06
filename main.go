package main

import (
	"github.com/ELTITO310/go-restapi/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
		AppName:       "RESTAPI GOLANG v0.0.1",
	})

	api := app.Group("/api")
	routes.Register(api)

	app.Listen(":3000")
}
