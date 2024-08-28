package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/kmdavidds/abdimasa-backend/internal/app/config"
)

// This is the boilerplate for vercel serverless functions

// Handler is the main entry point of the application. Think of it like the main() method
func Handler(w http.ResponseWriter, r *http.Request) {
	// This is needed to set the proper request path in `*fiber.Ctx`
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

// building the fiber application
func handler() http.HandlerFunc {
	db := config.NewDatabase()
	app := config.NewFiber()

	config.StartApp(&config.AppConfig{
		DB:  db,
		App: app,
	})

	return adaptor.FiberApp(app)
}
