package errors

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/validator"
)

var (
	ErrorNotFound = errors.New("record not found")
) 

func Handler(c *fiber.Ctx, err error) error {
	var e *validator.ValidationErrors
	if errors.As(err, &e) {
		return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
			"error": err,
		})
	}

	errorMappings := map[error]int{
		ErrorNotFound: http.StatusNotFound,
	}

	statusCode, exists := errorMappings[err]
	if !exists {
		statusCode = fiber.StatusInternalServerError
	}

	return c.Status(statusCode).JSON(map[string]any{
		"error": err.Error(),
	})
}