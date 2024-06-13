package routes

import (
	"cursogo/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/todos", controllers.GetTodos)
	api.Get("/todos/:id", controllers.GetTodoByID)
	api.Post("/todos", controllers.CreateTodo)
	api.Put("/todos/:id", controllers.UpdateTodoByID)
	api.Delete("/todos/:id", controllers.DeleteTodoByID)
}
