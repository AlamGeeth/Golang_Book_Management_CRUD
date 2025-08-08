package handlers

import (
	"book-management-app/database"
	"book-management-app/models"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	database.DB.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		return c.Status(404).SendString("Book not found")
	}
	return c.JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Create(&book)
	return c.Status(201).JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	result := database.DB.Delete(&models.Book{}, id)
	if result.RowsAffected == 0 {
		return c.Status(404).SendString("Book not found")
	}
	return c.SendString("Book deleted")
}
