package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kmdavidds/abdimasa-backend/internal/app/service"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/dto"
)

type AuthHandler interface {
	Login() fiber.Handler
}

type authHandler struct {
	as service.AuthService
}

func NewAuthHandler(
	as service.AuthService,
) AuthHandler {
	return &authHandler{as}
}

func (ah *authHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.LoginRequest{}
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		tokenString, err := ah.as.Login(req)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(map[string]any{
			"token": tokenString,
		})
	}
}