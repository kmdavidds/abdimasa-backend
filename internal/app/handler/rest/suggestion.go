package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kmdavidds/abdimasa-backend/internal/app/service"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/dto"
	"github.com/valyala/fasthttp"
)

type SuggestionHandler interface {
	Create() fiber.Handler
	GetAll() fiber.Handler
	Delete() fiber.Handler
}

type suggestionHandler struct {
	ss service.SuggestionService
}

func NewSuggestionHandler(
	ss service.SuggestionService,
) SuggestionHandler {
	return &suggestionHandler{ss}
}

func (sh *suggestionHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.CreateSuggestionRequest{}
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		attachmentFile1, err := c.FormFile("attachment1")
		if err != nil && err != fasthttp.ErrMissingFile {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": "failed to get attachment1",
			})
		}
		
		req.Attachment1 = attachmentFile1

		err = sh.ss.Create(req)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

func (sh *suggestionHandler) GetAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		suggestions, err := sh.ss.GetAll()
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(map[string]any{
			"suggestions": suggestions,
		})
	}
}

func (sh *suggestionHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.DeleteSuggestionRequest{}
		err := c.ParamsParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		err = sh.ss.Delete(req)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}
