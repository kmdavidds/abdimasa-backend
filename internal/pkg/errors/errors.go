package errors

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/validator"
)

var (
	ErrorNotFound           = errors.New("record not found")
	ErrorInvalidToken       = errors.New("invalid jwt token")
	ErrorExpiredToken       = errors.New("expired jwt token")
	ErrorInvalidCredentials = errors.New("invalid credentials")
)

func Handler(c *fiber.Ctx, err error) error {
	var e *validator.ValidationErrors
	if errors.As(err, &e) {
		return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
			"error": err,
		})
	}

	errorMappings := map[error]int{
		ErrorNotFound:           http.StatusNotFound,
		ErrorInvalidToken:       http.StatusUnauthorized,
		ErrorExpiredToken:       http.StatusUnauthorized,
		ErrorInvalidCredentials: http.StatusUnauthorized,
	}

	statusCode, exists := errorMappings[err]
	if !exists {
		statusCode = fiber.StatusInternalServerError
	}

	return c.Status(statusCode).JSON(map[string]any{
		"error": err.Error(),
	})
}
