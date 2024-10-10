package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kmdavidds/abdimasa-backend/internal/app/service"
	"github.com/kmdavidds/abdimasa-backend/internal/pkg/dto"
	"github.com/valyala/fasthttp"
)

type PlaceHandler interface {
	Create() fiber.Handler
	GetAll() fiber.Handler
	GetByID() fiber.Handler
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

		return c.Status(http.StatusOK).JSON(map[string]any{
			"places": places,
		})
	}
}

func (ph *placeHandler) GetByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.GetPlaceByIDRequest{}
		err := c.ParamsParser(&req)
		if err != nil {
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
				"error": err,
			})
		}

		place, err := ph.ps.GetByID(req)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(map[string]any{
			"place": place,
		})
	}
}

func (ph *placeHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.UpdatePlaceRequest{}
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
			return c.Status(http.StatusUnprocessableEntity).JSON(map[string]any{
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