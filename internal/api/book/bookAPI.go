package api

// import libraries
import (
	"example/database"
	"example/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CRUD API with book model

// Get all books
func GetAll(c *fiber.Ctx) error {
	db := database.DBConn
	var books []model.Book
	db.Find(&books)
	if len(books) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No books present", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Books Found", "data": books})
}

// Get single book based on id
func Get(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book model.Book
	db.Find(&book, "id = ?", id)

	if book.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No book present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Book Found", "data": book})
}

// Create an new book
func Create(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(model.Book)

	err := c.BodyParser(book)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	book.ID = uuid.New()
	err = db.Create(&book).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create book", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created Book", "data": book})
}

//todo: fix this api. This need to fetch body response and assign data to found book data.
// Update an book
func Update(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DBConn

	var book model.Book
	db.First(&book, "id = ?", id)
	if book.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No book present", "data": nil})
	}

	//db.Update(&book)

	return c.JSON(fiber.Map{"status": "success", "message": "book Found", "data": book})

}

// Delete an book
func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book model.Book
	db.First(&book, "id = ?", id)
	if book.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No book present", "data": nil})
	}
	err := db.Delete(&book, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete book", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Book"})
}
