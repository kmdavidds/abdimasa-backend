package config

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kmdavidds/abdimasa-backend/internal/app/handler/rest"
	"github.com/kmdavidds/abdimasa-backend/internal/app/handler/rest/middleware"
	"github.com/kmdavidds/abdimasa-backend/internal/app/handler/rest/route"
	"github.com/kmdavidds/abdimasa-backend/internal/app/repository"
	"github.com/kmdavidds/abdimasa-backend/internal/app/service"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/jwt"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/supabase"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/validator"
	"gorm.io/gorm"
)

type AppConfig struct {
	DB *gorm.DB
	App *fiber.App
}

func StartApp(config *AppConfig) {
	val := validator.NewValidator()
	jwt := jwt.New([]byte(os.Getenv("JWT_SECRET")), 24*time.Hour)
	supabase := supabase.New(
		os.Getenv("SUPABASE_PROJECT_URL"), 
		os.Getenv("SUPABASE_TOKEN"), 
		os.Getenv("SUPABASE_BUCKET_NAME"),
	)

	// repositories
	activityRepo := repository.NewActivityRepository(config.DB)
	placeRepo := repository.NewPlaceRepository(config.DB)
	businessRepo := repository.NewBusinessRepository(config.DB)
	remarkRepo := repository.NewRemarkRepository(config.DB)
	suggestionRepo := repository.NewSuggestionRepository(config.DB)
	newsRepo := repository.NewNewsRepository(config.DB)
	detailRepo := repository.NewDetailRepository(config.DB)

	// services
	activityService := service.NewActivityService(activityRepo, val, supabase)
	placeService := service.NewPlaceService(placeRepo, val, supabase)
	businessService := service.NewBusinessService(businessRepo, val, supabase)
	remarkService := service.NewRemarkService(remarkRepo, val)
	suggestionService := service.NewSuggestionService(suggestionRepo, val, supabase)
	newsService := service.NewNewsService(newsRepo, val, supabase)
	detailService := service.NewDetailService(detailRepo, val)
	authService := service.NewAuthService(val, jwt)

	// handlers
	activityHandler := rest.NewActivityHandler(activityService)
	placeHandler := rest.NewPlaceHandler(placeService)
	businessHandler := rest.NewBusinessHandler(businessService)
	remarkHandler := rest.NewRemarkHandler(remarkService)
	suggestionHandler := rest.NewSuggestionHandler(suggestionService)
	newsHandler := rest.NewNewsHandler(newsService)
	detailHandler := rest.NewDetailHandler(detailService)
	authHandler := rest.NewAuthHandler(authService)

	// middlewares
	authMiddleware := middleware.NewAuth(jwt)

	routes := route.Config{
		App: config.App,
		ActivityHandler: activityHandler,
		PlaceHandler: placeHandler,
		BusinessHandler: businessHandler,
		RemarkHandler: remarkHandler,
		SuggestionHandler: suggestionHandler,
		NewsHandler: newsHandler,
		DetailHandler: detailHandler,
		AuthHandler: authHandler,
		AuthMiddleware: authMiddleware,
	}

	routes.Register()
}