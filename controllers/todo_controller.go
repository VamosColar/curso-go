package controllers

import (
	"cursogo/config"
	"cursogo/models"
	"cursogo/views"
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	db := config.GetDB()
	rows, err := db.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		return c.Status(500).JSON(views.Response{Message: "Error fetching todos"})
	}
	defer rows.Close()

	todos := []models.Todo{}
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return c.Status(500).JSON(views.Response{Message: "Error scanning todo"})
		}
		todos = append(todos, todo)
	}
	return c.JSON(todos)
}

func GetTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := config.GetDB()
	row := db.QueryRow("SELECT id, title, completed FROM todos WHERE id = ?", id)

	var todo models.Todo
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(views.Response{Message: "Todo not found"})
		}
		return c.Status(500).JSON(views.Response{Message: "Error fetching todo"})
	}
	return c.JSON(todo)
}

func CreateTodo(c *fiber.Ctx) error {
	db := config.GetDB()
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(views.Response{Message: "Invalid input"})
	}

	result, err := db.Exec("INSERT INTO todos (title, completed) VALUES (?, ?)", todo.Title, todo.Completed)
	if err != nil {
		return c.Status(500).JSON(views.Response{Message: "Error creating todo"})
	}

	id, _ := result.LastInsertId()
	todo.ID = int(id)
	return c.Status(201).JSON(todo)
}

func UpdateTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := config.GetDB()
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).JSON(views.Response{Message: "Invalid input"})
	}

	_, err := db.Exec("UPDATE todos SET title = ?, completed = ? WHERE id = ?", todo.Title, todo.Completed, id)
	if err != nil {
		return c.Status(500).JSON(views.Response{Message: "Error updating todo"})
	}

	idInt, _ := strconv.Atoi(id)

	todo.ID = idInt

	return c.JSON(todo)
}

func DeleteTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := config.GetDB()

	_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return c.Status(500).JSON(views.Response{Message: "Error deleting todo"})
	}
	return c.JSON(views.Response{Message: "Todo deleted"})
}
