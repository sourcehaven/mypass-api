package routes

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ParseBody(c *fiber.Ctx, obj any) error {
	err := c.BodyParser(obj)
	if err != nil {
		SendResponse(c, http.StatusBadRequest, "Error while parsing body", nil)
		return err
	}
	return nil
}
