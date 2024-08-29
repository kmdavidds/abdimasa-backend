package config

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/errors"
)

func NewFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: errors.Handler,
	})

	app.Use(cors.New())

	return app
}
