package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kmdavidds/abdimasa-backend/internal/app/handler/rest"
	"github.com/kmdavidds/abdimasa-backend/internal/app/handler/rest/middleware"
)

type Config struct {
	App               *fiber.App
	ActivityHandler   rest.ActivityHandler
	PlaceHandler      rest.PlaceHandler
	BusinessHandler   rest.BusinessHandler
	RemarkHandler     rest.RemarkHandler
	SuggestionHandler rest.SuggestionHandler
	NewsHandler       rest.NewsHandler
	DetailHandler     rest.DetailHandler
	AuthHandler       rest.AuthHandler
	AuthMiddleware    middleware.Auth
}

func (c *Config) Register() {
	api := c.App.Group("/api")

	v1 := api.Group("/v1")

	c.activityRoutes(v1)
	c.placeRoutes(v1)
	c.businessRoutes(v1)
	c.remarkRoutes(v1)
	c.suggestionRoutes(v1)
	c.newsRoutes(v1)
	c.detailRoutes(v1)
	c.authRoutes(v1)
}

func (c *Config) activityRoutes(r fiber.Router) {
	activities := r.Group("/activities")

	activities.Get("", c.ActivityHandler.GetAll())

	activities.Post("", c.AuthMiddleware.Authenticate(), c.ActivityHandler.Create())
	activities.Put("", c.AuthMiddleware.Authenticate(),  c.ActivityHandler.Update())
	activities.Delete("/:id", c.AuthMiddleware.Authenticate(),  c.ActivityHandler.Delete())
}

func (c *Config) placeRoutes(r fiber.Router) {
	places := r.Group("/places")

	places.Get("", c.PlaceHandler.GetAll())
	places.Get("/:id", c.PlaceHandler.GetByID())

	places.Post("", c.AuthMiddleware.Authenticate(),  c.PlaceHandler.Create())
	places.Put("", c.AuthMiddleware.Authenticate(),  c.PlaceHandler.Update())
	places.Delete("/:id", c.AuthMiddleware.Authenticate(),  c.PlaceHandler.Delete())
}

func (c *Config) businessRoutes(r fiber.Router) {
	businesses := r.Group("/businesses")

	businesses.Get("", c.BusinessHandler.GetAll())
	businesses.Get("/:id", c.BusinessHandler.GetByID())

	businesses.Post("", c.AuthMiddleware.Authenticate(),  c.BusinessHandler.Create())
	businesses.Put("", c.AuthMiddleware.Authenticate(),  c.BusinessHandler.Update())
	businesses.Delete("/:id", c.AuthMiddleware.Authenticate(),  c.BusinessHandler.Delete())
}

func (c *Config) remarkRoutes(r fiber.Router) {
	remarks := r.Group("/remarks")

	remarks.Get("", c.RemarkHandler.GetAll())

	remarks.Post("", c.AuthMiddleware.Authenticate(),  c.RemarkHandler.Create())
}

func (c *Config) suggestionRoutes(r fiber.Router) {
	suggestions := r.Group("/suggestions")

	suggestions.Get("", c.SuggestionHandler.GetAll())

	suggestions.Post("", c.AuthMiddleware.Authenticate(),  c.SuggestionHandler.Create())
	suggestions.Delete("/:id", c.AuthMiddleware.Authenticate(),  c.SuggestionHandler.Delete())
}

func (c *Config) newsRoutes(r fiber.Router) {
	news := r.Group("/news")

	news.Get("", c.NewsHandler.GetAll())
	news.Get("/:id", c.NewsHandler.GetByID())

	news.Post("", c.AuthMiddleware.Authenticate(),  c.NewsHandler.Create())
	news.Put("", c.AuthMiddleware.Authenticate(),  c.NewsHandler.Update())
	news.Delete("/:id", c.AuthMiddleware.Authenticate(),  c.NewsHandler.Delete())
}

func (c *Config) detailRoutes(r fiber.Router) {
	detail := r.Group("/details")

	detail.Get("", c.DetailHandler.GetAll())
	detail.Get("/:slug", c.DetailHandler.GetBySlug())

	detail.Post("", c.AuthMiddleware.Authenticate(),  c.DetailHandler.Create())
	detail.Put("", c.AuthMiddleware.Authenticate(),  c.DetailHandler.Update())
	detail.Delete("/:id", c.AuthMiddleware.Authenticate(),  c.DetailHandler.Delete())
}

func (c *Config) authRoutes(r fiber.Router) {
	auth := r.Group("/auth")

	auth.Post("/login", c.AuthHandler.Login())
}
