package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sourcehaven/mypass-api/internal/model"
	"net/http"
	"time"
)

// NewResponse creates a new Response object with a message and optional data
func NewResponse(
	stat string,
	msg string,
	data any,
) model.Response {
	resp := model.Response{
		Status:    stat,
		Message:   msg,
		Data:      data,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		RequestId: uuid.NewString(),
	}
	return resp
}

func SendResponse(
	c *fiber.Ctx,
	code int,
	msg string,
	data any,
) {
	c.Status(code).JSON(
		NewResponse(
			http.StatusText(code),
			msg,
			data,
		))
}
