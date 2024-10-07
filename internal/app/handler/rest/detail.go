package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kmdavidds/abdimasa-backend/internal/app/service"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/dto"
)

type DetailHandler interface {
	Create() fiber.Handler
	GetAll() fiber.Handler
	GetByID() fiber.Handler
	GetBySlug() fiber.Handler
	Update() fiber.Handler
	Delete() fiber.Handler
}

type detailHandler struct {
	ds service.DetailService
}

func NewDetailHandler(
	ds service.DetailService,
) DetailHandler {
	return &detailHandler{ds}
}

func (dh *detailHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.CreateDetailRequest{}
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		err = dh.ds.Create(req)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

func (dh *detailHandler) GetAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		detail, err := dh.ds.GetAll()
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(map[string]any{
			"detail": detail,
		})
	}
}

func (dh *detailHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.UpdateDetailRequest{}
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		err = dh.ds.Update(req)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

func (dh *detailHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.DeleteDetailRequest{}
		err := c.ParamsParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		err = dh.ds.Delete(req)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

func (dh *detailHandler) GetByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.GetDetailByIDRequest{}
		err := c.ParamsParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		detail, err := dh.ds.GetByID(req)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(map[string]any{
			"detail": detail,
		})
	}
}

func (dh *detailHandler) GetBySlug() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.GetDetailBySlugRequest{}
		err := c.ParamsParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		detail, err := dh.ds.GetBySlug(req)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(map[string]any{
			"detail": detail,
		})
	}
}
