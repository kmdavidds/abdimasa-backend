package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kmdavidds/abdimasa-backend/internal/app/service"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/dto"
)

type RemarkHandler interface {
	Create() fiber.Handler
	GetAll() fiber.Handler
}

type remarkHandler struct {
	rs service.RemarkService
}

func NewRemarkHandler(
	rs service.RemarkService,
) RemarkHandler {
	return &remarkHandler{rs}
}

func (rh *remarkHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.CreateRemarkRequest{}
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		err = rh.rs.Create(req)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

func (rh *remarkHandler) GetAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		remarks, err := rh.rs.GetAll()
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(map[string]any{
			"remarks": remarks,
		})
	}
}