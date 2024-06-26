package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nazzarr03/TO-DO-Proxolab/controller"
)

func main() {
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})

	app.Post("/todo", controller.CreateTodo)
	app.Get("/todos", controller.GetTodos)

	app.Listen(":8081")
}
