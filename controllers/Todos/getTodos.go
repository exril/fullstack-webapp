package controllers

import (
	"context"

	"github.com/ftpskid/fullstack-webapp/initializers"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	initializers.LoadMongo()
}

var collection *mongo.Collection

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

func GetTodos(c *fiber.Ctx) error {
	var todos []Todo

	//define cursor here and use it after deff err
	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
	}

	return c.JSON(todos)
}
