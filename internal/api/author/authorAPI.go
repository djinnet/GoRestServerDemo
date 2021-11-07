package api

// import libraries
import (
	"example/database"
	"example/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CRUD API with author model

// Get all authors
func GetAll(c *fiber.Ctx) error {
	db := database.DBConn
	var authors []model.Author
	db.Find(&authors)
	if len(authors) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No authors present", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "authors Found", "data": authors})
}

// Get single author based on id
func Get(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var author model.Author
	db.Find(&author, "id = ?", id)

	if author.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No author present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "author Found", "data": author})
}

// Create an new author
func Create(c *fiber.Ctx) error {
	db := database.DBConn
	author := new(model.Author)

	err := c.BodyParser(author)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	author.ID = uuid.New()
	err = db.Create(&author).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create author", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created author", "data": author})
}

//todo: fix this api. This need to fetch body response and assign data to found book data.
// Update an author
func Update(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DBConn

	var author model.Author
	db.First(&author, "id = ?", id)
	if author.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No author present", "data": nil})
	}

	//db.Update(&author)

	return c.JSON(fiber.Map{"status": "success", "message": "Author Found", "data": author})

}

// Delete an author
func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var author model.Author
	db.First(&author, "id = ?", id)
	if author.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No author present", "data": nil})
	}
	err := db.Delete(&author, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete author", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted Author"})
}
