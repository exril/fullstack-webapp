package main

import (
	"fmt"
	"log"
	"os"

	controllers "github.com/ftpskid/fullstack-webapp/controllers/Todos"
	"github.com/ftpskid/fullstack-webapp/initializers"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnvInitializers()
	initializers.LoadMongo()
}

func main() {
	fmt.Printf("Connecting..\n")

	app := fiber.New()

	// Build TODOS here
	app.Post("/api/user/login", controllers.UserLogin)
	app.Post("/api/user/signup", controllers.UserSignup)
	app.Get("/api/todos", controllers.CreateTodos)
	app.Post("/api/todos", controllers.CreateTodos)
	app.Patch("/api/todos/:id", controllers.UpdateTodo)
	app.Delete("/api/todos/:id", controllers.DeleteTodos)

	// Build Timer here

	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))

}
