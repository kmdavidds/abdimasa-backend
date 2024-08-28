package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kmdavidds/abdimasa-backend/internal/app/handler/rest"
)

type Config struct {
	App             *fiber.App
	ActivityHandler rest.ActivityHandler
	PlaceHandler    rest.PlaceHandler
	BusinessHandler rest.BusinessHandler
}

func (c *Config) Register() {
	api := c.App.Group("/api")

	v1 := api.Group("/v1")

	c.activityRoutes(v1)
	c.placeRoutes(v1)
	c.businessRoutes(v1)
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
	places.Put("", c.PlaceHandler.Update())
	places.Delete("/:id", c.PlaceHandler.Delete())
}

func (c *Config) businessRoutes(r fiber.Router) {
	businesses := r.Group("/businesses")

	businesses.Post("", c.BusinessHandler.Create())
	businesses.Get("", c.BusinessHandler.GetAll())
	businesses.Put("", c.BusinessHandler.Update())
	businesses.Delete("/:id", c.BusinessHandler.Delete())
}