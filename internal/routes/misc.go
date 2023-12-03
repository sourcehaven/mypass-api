package routes

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// Teapot sends back HTTP 418 I'm a teapot status.
// @Summary I'm a teapot
// @Description Responds with HTTP status 418 I'm a teapot.
// @Tags examples
// @Produce string
// @Success 418 {string} string "I am a teapot!"
// @Router /teapot [get]
func Teapot(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusTeapot).SendString("I am a teapot!")
}
