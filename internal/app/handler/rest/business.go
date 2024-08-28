package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kmdavidds/abdimasa-backend/internal/app/service"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/dto"
)

type BusinessHandler interface {
	Create() fiber.Handler
	GetAll() fiber.Handler
	Update() fiber.Handler
	Delete() fiber.Handler
}

type businessHandler struct {
	bs service.BusinessService
}

func NewBusinessHandler(
	bs service.BusinessService,
) BusinessHandler {
	return &businessHandler{bs}
}

func (bh *businessHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.CreateBusinessRequest{}
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]interface{}{
				"error": err,
			})
		}

		err = bh.bs.Create(req)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

func (bh *businessHandler) GetAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		businesses, err := bh.bs.GetAll()
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"businesses": businesses,
		})
	}
}

func (bh *businessHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.UpdateBusinessRequest{}
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]interface{}{
				"error": err,
			})
		}

		err = bh.bs.Update(req)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}

func (bh *businessHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.DeleteBusinessRequest{}
		err := c.ParamsParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]interface{}{
				"error": err,
			})
		}

		err = bh.bs.Delete(req)
		if err != nil {
			return err
		}

		return c.SendStatus(http.StatusNoContent)
	}
}