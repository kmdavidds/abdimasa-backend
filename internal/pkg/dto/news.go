package dto

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type CreateNewsRequest struct {
	Title       string                `form:"title" validate:"required,min=1,max=255"`
	Description string                `form:"description" validate:"required"`
	Image1      *multipart.FileHeader `form:"image1"`
}

type GetNewsByIDRequest struct {
	ID uuid.UUID `param:"id" validate:"required"`
}

type UpdateNewsRequest struct {
	ID          uuid.UUID             `form:"id" validate:"required"`
	Title       string                `form:"title" validate:"max=255"`
	Description string                `form:"description"`
	Image1      *multipart.FileHeader `form:"image1"`
}

type DeleteNewsRequest struct {
	ID uuid.UUID `param:"id" validate:"required"`
}
