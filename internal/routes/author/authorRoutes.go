package authorRoutes

import (
	Api "example/internal/api/author"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {
	author := router.Group("/author")
	author.Get("/", Api.GetAll)
	author.Get("/:id", Api.Get)
	author.Put("/:id", Api.Update)
	author.Post("/", Api.Create)
	author.Delete("/:id", Api.Delete)
}
