package vercel

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kmdavidds/abdimasa-backend/internal/app/config"
)

// This is the same as cmd/app/main.go
// This avoids the error of vercel importing a package named "config"

func NewVercelApp() *fiber.App {
	db := config.NewDatabase()
	app := config.NewFiber()

	config.StartApp(&config.AppConfig{
		DB:  db,
		App: app,
	})
	
	app.Use(cors.New())

	return app
}