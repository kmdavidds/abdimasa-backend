package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kmdavidds/abdimasa-backend/internal/app/service"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/dto"
)

type NewsHandler interface {
	Create() fiber.Handler
	GetAll() fiber.Handler
	GetByID() fiber.Handler
	Update() fiber.Handler
	Delete() fiber.Handler
}

type newsHandler struct {
	ns service.NewsService
}

func NewNewsHandler(
	ns service.NewsService,
) NewsHandler {
	return &newsHandler{ns}
}

func (nh *newsHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.CreateNewsRequest{}
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		err = nh.ns.Create(req)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

func (nh *newsHandler) GetAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		news, err := nh.ns.GetAll()
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(map[string]any{
			"news": news,
		})
	}
}

func (nh *newsHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.UpdateNewsRequest{}
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		err = nh.ns.Update(req)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

func (nh *newsHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.DeleteNewsRequest{}
		err := c.ParamsParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		err = nh.ns.Delete(req)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

func (nh *newsHandler) GetByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.GetNewsByIDRequest{}
		err := c.ParamsParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		news, err := nh.ns.GetByID(req)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(map[string]any{
			"news": news,
		})
	}
}
