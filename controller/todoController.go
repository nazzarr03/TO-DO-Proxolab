package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nazzarr03/TO-DO-Proxolab/database"
	"github.com/nazzarr03/TO-DO-Proxolab/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var todoCollection = database.OpenCollection(database.Client, "todo")

func CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
			"data":  nil,
		})
	}

	if todo.Description == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Description is required",
			"data":  nil,
		})
	}

	todo.ID = primitive.NewObjectID()
	todo.Completed = false
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := todoCollection.InsertOne(ctx, todo)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"data":  nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Todo created successfully",
		"data":    todo,
	})
}

func GetTodos(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := todoCollection.Find(ctx, primitive.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"data":  nil,
		})
	}
	defer cursor.Close(ctx)

	var todos []models.Todo
	if err = cursor.All(ctx, &todos); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"data":  nil,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": todos,
	})
}

func UpdateTodo(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
			"data":  nil,
		})
	}

	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
			"data":  nil,
		})
	}

	update := primitive.D{{Key: "$set", Value: todo}}
	err = todoCollection.FindOneAndUpdate(ctx, primitive.M{"_id": objID}, update).Err()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"data":  nil,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Todo updated successfully",
	})
}
