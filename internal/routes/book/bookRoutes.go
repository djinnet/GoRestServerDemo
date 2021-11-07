package bookRoutes

import (
	Api "example/internal/api/book"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	book := router.Group("/book")
	book.Get("/", Api.GetAll)
	book.Get("/:id", Api.Get)
	book.Put("/:id", Api.Update)
	book.Post("/", Api.Create)
	book.Delete("/:id", Api.Delete)
}
