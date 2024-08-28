package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kmdavidds/abdimasa-backend/internal/app/handler/rest"
	"github.com/kmdavidds/abdimasa-backend/internal/app/handler/rest/route"
	"github.com/kmdavidds/abdimasa-backend/internal/app/repository"
	"github.com/kmdavidds/abdimasa-backend/internal/app/service"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/validator"
	"gorm.io/gorm"
)

type AppConfig struct {
	DB *gorm.DB
	App *fiber.App
}

func StartApp(config *AppConfig) {
	val := validator.NewValidator()

	// repositories
	activityRepo := repository.NewActivityRepository(config.DB)

	// services
	activityService := service.NewActivityService(activityRepo, val)

	// handlers
	activityHandler := rest.NewActivityHandler(activityService)

	routes := route.Config{
		App: config.App,
		ActivityHandler: activityHandler,
	}

	routes.Register()
}