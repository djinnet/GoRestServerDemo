package errorHandler

import (
	"github.com/gofiber/fiber/v2"
)

// HTTPError defines the field which should be the response.
type HTTPError struct {
	Status string `json:"status" example:"error"`          // The status error of string.
	Error  string `json:"error" example:"a message error"` // The message error description of string.
}

// HTTPSuccess defines the field which should be the response.
type HTTPSuccess struct {
	Status string      `json:"status" example:"success"` // The status success of string.
	Data   interface{} `json:"data"`                     // The result of json.
}

// ErrorHandler is used to catch error thrown inside the routes by ctx.Next(err)
func ErrorHandler(c *fiber.Ctx, err error) error {
	// Statuscode defaults to 500
	code := fiber.StatusInternalServerError

	// Check if it's an fiber.Error type
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(&HTTPError{
		Status: "error",
		Error:  err.Error(),
	})
}
