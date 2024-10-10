package dto

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type CreatePlaceRequest struct {
	Name         string                `json:"name" validate:"required,min=1,max=255"`
	Location     string                `json:"location" validate:"required,min=1,max=255"`
	Description  string                `json:"description" validate:"required,min=1"`
	Address      string                `json:"address" validate:"required,min=1,max=255"`
	OpeningHours string                `json:"openingHours" validate:"min=9,max=10"`
	ClosingHours string                `json:"closingHours" validate:"min=9,max=10"`
	EntryPrice   string                `json:"entryPrice" validate:"min=1,max=255"`
	Image1       *multipart.FileHeader `form:"image1"`
	Image2       *multipart.FileHeader `form:"image2"`
	Image3       *multipart.FileHeader `form:"image3"`
	MapURL       string                `json:"mapURL" validate:"url"`
	Rating       int                   `json:"rating" validate:"min=0,max=50"`
}

type GetPlaceByIDRequest struct {
	ID uuid.UUID `param:"id" validate:"required"`
}

type UpdatePlaceRequest struct {
	ID           uuid.UUID             `json:"id" validate:"required"`
	Name         string                `json:"name" validate:"max=255"`
	Location     string                `json:"location" validate:"max=255"`
	Description  string                `json:"description" validate:""`
	Address      string                `json:"address" validate:"max=255"`
	OpeningHours string                `json:"openingHours" validate:"max=10"`
	ClosingHours string                `json:"closingHours" validate:"max=10"`
	EntryPrice   string                `json:"entryPrice" validate:"max=255"`
	Image1       *multipart.FileHeader `form:"image1"`
	Image2       *multipart.FileHeader `form:"image2"`
	Image3       *multipart.FileHeader `form:"image3"`
	MapURL       string                `json:"mapURL" validate:"url"`
	Rating       int                   `json:"rating" validate:"max=50"`
}

type DeletePlaceRequest struct {
	ID uuid.UUID `param:"id" validate:"required"`
}
