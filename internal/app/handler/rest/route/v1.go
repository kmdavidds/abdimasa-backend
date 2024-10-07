package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kmdavidds/abdimasa-backend/internal/app/handler/rest"
)

type Config struct {
	App               *fiber.App
	ActivityHandler   rest.ActivityHandler
	PlaceHandler      rest.PlaceHandler
	BusinessHandler   rest.BusinessHandler
	RemarkHandler     rest.RemarkHandler
	SuggestionHandler rest.SuggestionHandler
	NewsHandler       rest.NewsHandler
	DetailHandler       rest.DetailHandler
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
}

func (c *Config) activityRoutes(r fiber.Router) {
	activities := r.Group("/activities")

	activities.Post("", c.ActivityHandler.Create())
	activities.Get("", c.ActivityHandler.GetAll())
	activities.Put("", c.ActivityHandler.Update())
	activities.Delete("/:id", c.ActivityHandler.Delete())
}

func (c *Config) placeRoutes(r fiber.Router) {
	places := r.Group("/places")

	places.Post("", c.PlaceHandler.Create())
	places.Get("", c.PlaceHandler.GetAll())
	places.Get("/:id", c.PlaceHandler.GetByID())
	places.Put("", c.PlaceHandler.Update())
	places.Delete("/:id", c.PlaceHandler.Delete())
}

func (c *Config) businessRoutes(r fiber.Router) {
	businesses := r.Group("/businesses")

	businesses.Post("", c.BusinessHandler.Create())
	businesses.Get("", c.BusinessHandler.GetAll())
	businesses.Get("/:id", c.BusinessHandler.GetByID())
	businesses.Put("", c.BusinessHandler.Update())
	businesses.Delete("/:id", c.BusinessHandler.Delete())
}

func (c *Config) remarkRoutes(r fiber.Router) {
	remarks := r.Group("/remarks")

	remarks.Post("", c.RemarkHandler.Create())
	remarks.Get("", c.RemarkHandler.GetAll())
}

func (c *Config) suggestionRoutes(r fiber.Router) {
	suggestions := r.Group("/suggestions")

	suggestions.Post("", c.SuggestionHandler.Create())
	suggestions.Get("", c.SuggestionHandler.GetAll())
	suggestions.Delete("/:id", c.SuggestionHandler.Delete())
}

func (c *Config) newsRoutes(r fiber.Router) {
	news := r.Group("/news")

	news.Post("", c.NewsHandler.Create())
	news.Get("", c.NewsHandler.GetAll())
	news.Get("/:id", c.NewsHandler.GetByID())
	news.Put("", c.NewsHandler.Update())
	news.Delete("/:id", c.NewsHandler.Delete())
}

func (c *Config) detailRoutes(r fiber.Router) {
	detail := r.Group("/details")

	detail.Post("", c.DetailHandler.Create())
	detail.Get("", c.DetailHandler.GetAll())
	detail.Get("/:slug", c.DetailHandler.GetBySlug())
	detail.Put("", c.DetailHandler.Update())
	detail.Delete("/:id", c.DetailHandler.Delete())
}