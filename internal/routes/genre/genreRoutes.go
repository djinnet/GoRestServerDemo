package genreRoutes

import (
	Handler "example/internal/api/genre"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	genre := router.Group("/genre")
	genre.Get("/", Handler.GetAll)
	genre.Get("/:id", Handler.Get)
	genre.Put("/:id", Handler.Update)
	genre.Post("/", Handler.Create)
	genre.Delete("/:id", Handler.Delete)
}
