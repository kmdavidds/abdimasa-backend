package dto

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type CreateBusinessRequest struct {
	Name        string                `form:"name" validate:"required,min=1,max=255"`
	Location    string                `form:"location" validate:"required,min=1,max=255"`
	Description string                `form:"description" validate:"required,min=1"`
	Address     string                `form:"address" validate:"required,min=1,max=255"`
	PriceRange  string                `form:"priceRange" validate:"min=1,max=255"`
	Image1      *multipart.FileHeader `form:"image1"`
	Image2      *multipart.FileHeader `form:"image2"`
	Image3      *multipart.FileHeader `form:"image3"`
	Contact     string                `form:"contact" validate:"min=1,max=255"`
	MapURL      string                `form:"mapURL" validate:"url"`
	Rating      int                   `form:"rating" validate:"min=0,max=50"`
}

type GetBusinessByIDRequest struct {
	ID uuid.UUID `param:"id" validate:"required"`
}

type UpdateBusinessRequest struct {
	ID          uuid.UUID             `form:"id" validate:"required"`
	Name        string                `form:"name" validate:"max=255"`
	Location    string                `form:"location" validate:"max=255"`
	Description string                `form:"description" validate:""`
	Address     string                `form:"address" validate:"max=255"`
	PriceRange  string                `form:"priceRange" validate:"max=255"`
	Image1      *multipart.FileHeader `form:"image1"`
	Image2      *multipart.FileHeader `form:"image2"`
	Image3      *multipart.FileHeader `form:"image3"`
	Contact     string                `form:"contact" validate:"max=255"`
	MapURL      string                `form:"mapURL" validate:"url"`
	Rating      int                   `form:"rating" validate:"max=50"`
}

type DeleteBusinessRequest struct {
	ID uuid.UUID `param:"id" validate:"required"`
}
