package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id) // primitive usage as the id is in string and i want to use it as int

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid"})
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"compeleted": true}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}