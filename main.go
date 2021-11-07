package main

// import the libraries
import (
	"example/database"
	handler "example/internal/handler/errors"
	"example/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title Book Management System - BMS
// @version 1.0
// @description This is a sample swagger for Fiber

// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
func main() {
	// use fiber as application with an errorhandler
	app := fiber.New(fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	})

	// init database
	database.InitDatabase()

	// add Logger middleware with config
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// add cors middleware
	app.Use(cors.New())

	// add recover middleware
	app.Use(recover.New())

	// setup the routes
	router.SetupRoutes(app)

	//listen to port 3000 and check for error
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
