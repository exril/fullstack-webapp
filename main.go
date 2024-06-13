package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ftpskid/fullstack-webapp/controllers"
	"github.com/ftpskid/fullstack-webapp/initializers"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ID = ""

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

func init() {
	initializers.LoadEnvInitializers()
	initializers.LoadMongo()
}

func main() {
	fmt.Printf("Connecting..\n")

	app := fiber.New()

	app.Get("/api/todos", controllers.GetTodos)
	app.Post("/api/todos", controllers.CreateTodos)
	app.Patch("/api/todos/:id", controllers.UpdateTodo)
	app.Delete("/api/todos/:id", controllers.DeleteTodos)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))

}
