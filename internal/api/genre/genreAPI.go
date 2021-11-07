package api

import (
	"example/database"
	"example/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAll(c *fiber.Ctx) error {
	db := database.DBConn
	var genres []model.Genre
	db.Find(&genres)
	if len(genres) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No genres present", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "genres Found", "data": genres})
}

func Get(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var genre model.Genre
	db.Find(&genre, "id = ?", id)

	if genre.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No genre present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "genre Found", "data": genre})
}

func Create(c *fiber.Ctx) error {
	db := database.DBConn
	genre := new(model.Genre)

	err := c.BodyParser(genre)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	genre.ID = uuid.New()
	err = db.Create(&genre).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create genre", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created genre", "data": genre})
}

//todo: fix this api. This need to fetch body response and assign data to found genre data.
func Update(c *fiber.Ctx) error {
	id := c.Params("id")

	db := database.DBConn

	var genre model.Genre
	db.First(&genre, "id = ?", id)
	if genre.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No genre present", "data": nil})
	}

	//db.Update(&genre)

	return c.JSON(fiber.Map{"status": "success", "message": "genre Found", "data": genre})

}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var genre model.Genre
	db.First(&genre, "id = ?", id)
	if genre.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No genre present", "data": nil})
	}
	err := db.Delete(&genre, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete genre", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted genre"})
}
