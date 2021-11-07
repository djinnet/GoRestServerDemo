package router

// import the libraries
import (
	swagger "github.com/arsmn/fiber-swagger/v2"

	authorRoutes "example/internal/routes/author"
	bookRoutes "example/internal/routes/book"
	genreRoutes "example/internal/routes/genre"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {

	// get swagger
	app.Get("/swagger/*", swagger.Handler)

	// group application in api
	api := app.Group("api", logger.New())

	// group api in v1
	v1 := api.Group("/v1")

	// setup the routes
	bookRoutes.SetupRoutes(v1)
	authorRoutes.SetupRoutes(v1)
	genreRoutes.SetupRoutes(v1)
}
