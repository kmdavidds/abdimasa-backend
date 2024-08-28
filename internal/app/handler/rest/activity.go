package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kmdavidds/abdimasa-backend/internal/app/service"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/dto"
)

type ActivityHandler interface {
	Create() fiber.Handler
	GetAll() fiber.Handler
	Update() fiber.Handler
	Delete() fiber.Handler
}

type activityHandler struct {
	as service.ActivityService
}

func NewActivityHandler(
	as service.ActivityService,
) ActivityHandler {
	return &activityHandler{as}
}

func (ah *activityHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.CreateActivityRequest{}
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		err = ah.as.Create(req)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

func (ah *activityHandler) GetAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		activities, err := ah.as.GetAll()
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(map[string]any{
			"activities": activities,
		})
	}
}

func (ah *activityHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.UpdateActivityRequest{}
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		err = ah.as.Update(req)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

func (ah *activityHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.DeleteActivityRequest{}
		err := c.ParamsParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		err = ah.as.Delete(req)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}