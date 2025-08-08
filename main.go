package main

import (
	"book-management-app/database"
	"book-management-app/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	app.Get("/books", handlers.GetBooks)
	app.Get("/books/:id", handlers.GetBook)
	app.Post("/books", handlers.CreateBook)
	app.Delete("/books/:id", handlers.DeleteBook)

	app.Listen(":8080")
}
