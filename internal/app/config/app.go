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
	placeRepo := repository.NewPlaceRepository(config.DB)
	businessRepo := repository.NewBusinessRepository(config.DB)
	remarkRepo := repository.NewRemarkRepository(config.DB)
	suggestionRepo := repository.NewSuggestionRepository(config.DB)

	// services
	activityService := service.NewActivityService(activityRepo, val)
	placeService := service.NewPlaceService(placeRepo, val)
	businessService := service.NewBusinessService(businessRepo, val)
	remarkService := service.NewRemarkService(remarkRepo, val)
	suggestionService := service.NewSuggestionService(suggestionRepo, val)

	// handlers
	activityHandler := rest.NewActivityHandler(activityService)
	placeHandler := rest.NewPlaceHandler(placeService)
	businessHandler := rest.NewBusinessHandler(businessService)
	remarkHandler := rest.NewRemarkHandler(remarkService)
	suggestionHandler := rest.NewSuggestionHandler(suggestionService)

	routes := route.Config{
		App: config.App,
		ActivityHandler: activityHandler,
		PlaceHandler: placeHandler,
		BusinessHandler: businessHandler,
		RemarkHandler: remarkHandler,
		SuggestionHandler: suggestionHandler,
	}

	routes.Register()
}