package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kmdavidds/abdimasa-backend/internal/app/service"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/dto"
)

type PlaceHandler interface {
	Create() fiber.Handler
	GetAll() fiber.Handler
	Update() fiber.Handler
	Delete() fiber.Handler
}

type placeHandler struct {
	ps service.PlaceService
}

func NewPlaceHandler(
	ps service.PlaceService,
) PlaceHandler {
	return &placeHandler{ps}
}

func (ph *placeHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.CreatePlaceRequest{}
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]interface{}{
				"error": err,
			})
		}

		err = ph.ps.Create(req)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

func (ph *placeHandler) GetAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		places, err := ph.ps.GetAll()
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"places": places,
		})
	}
}

func (ph *placeHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.UpdatePlaceRequest{}
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]interface{}{
				"error": err,
			})
		}

		err = ph.ps.Update(req)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

func (ph *placeHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.DeletePlaceRequest{}
		err := c.ParamsParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]interface{}{
				"error": err,
			})
		}

		err = ph.ps.Delete(req)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}