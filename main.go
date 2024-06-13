package main

import (
	"cursogo/config"
	"cursogo/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.InitDatabase()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
