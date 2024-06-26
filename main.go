package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nazzarr03/TO-DO-Proxolab/controller"
)

func main() {
	app := fiber.New()

	app.Post("/todo", controller.CreateTodo)
	app.Get("/todos", controller.GetTodos)
	app.Put("/todo/:id", controller.UpdateTodo)
	app.Delete("/todo/:id", controller.DeleteTodo)

	app.Listen(":8081")
}
