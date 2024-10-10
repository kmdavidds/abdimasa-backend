package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kmdavidds/abdimasa-backend/internal/app/service"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/dto"
	"github.com/valyala/fasthttp"
)

type BusinessHandler interface {
	Create() fiber.Handler
	GetAll() fiber.Handler
	GetByID() fiber.Handler
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
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		imageFile1, err := c.FormFile("image1")
		if err != nil && err != fasthttp.ErrMissingFile {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": "failed to get image1",
			})
		}
		imageFile2, err := c.FormFile("image2")
		if err != nil && err != fasthttp.ErrMissingFile {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": "failed to get image2",
			})
		}
		imageFile3, err := c.FormFile("image3")
		if err != nil && err != fasthttp.ErrMissingFile {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": "failed to get image3",
			})
		}
		
		req.Image1 = imageFile1
		req.Image2 = imageFile2
		req.Image3 = imageFile3

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

		return c.Status(http.StatusOK).JSON(map[string]any{
			"businesses": businesses,
		})
	}
}

func (bh *businessHandler) GetByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.GetBusinessByIDRequest{}
		err := c.ParamsParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		business, err := bh.bs.GetByID(req)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(map[string]any{
			"business": business,
		})
	}
}

func (bh *businessHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.UpdateBusinessRequest{}
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		imageFile1, err := c.FormFile("image1")
		if err != nil && err != fasthttp.ErrMissingFile {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": "failed to get image1",
			})
		}
		imageFile2, err := c.FormFile("image2")
		if err != nil && err != fasthttp.ErrMissingFile {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": "failed to get image2",
			})
		}
		imageFile3, err := c.FormFile("image3")
		if err != nil && err != fasthttp.ErrMissingFile {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": "failed to get image3",
			})
		}
		
		req.Image1 = imageFile1
		req.Image2 = imageFile2
		req.Image3 = imageFile3

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
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
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
