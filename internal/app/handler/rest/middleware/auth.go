package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/errors"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/jwt"
)

type Auth interface {
	Authenticate() fiber.Handler
}

type auth struct {
	jwt jwt.JWT
}

func NewAuth(
	jwt jwt.JWT,
) Auth {
	return &auth{jwt: jwt}
}

func (a *auth) Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		bearer := c.Get("Authorization")
		if bearer == "" {
			return errors.ErrorInvalidToken
		}

		tokenSlice := strings.Split(bearer, " ")
		if len(tokenSlice) != 2 {
			return errors.ErrorInvalidToken
		}

		tokenString := tokenSlice[1]
		
		err := a.jwt.Parse(tokenString)
		if err != nil {
			return err
		}
		return c.Next()
	}
}