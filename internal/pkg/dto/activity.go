package dto

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type CreateActivityRequest struct {
	Title    string                `form:"title" validate:"required,min=1,max=255"`
	Image1   *multipart.FileHeader `form:"image1"`
	Date     string                `form:"date" validate:"required,len=10"`
	Time     string                `form:"time" validate:"min=9,max=10"`
	Location string                `form:"location" validate:"required,min=1,max=255"`
}

type UpdateActivityRequest struct {
	ID       uuid.UUID             `form:"id"`
	Title    string                `form:"title" validate:"max=255"`
	Image1   *multipart.FileHeader `form:"image1"`
	Date     string                `form:"date"`
	Time     string                `form:"time" validate:"max=10"`
	Location string                `form:"location" validate:"max=255"`
}

type DeleteActivityRequest struct {
	ID uuid.UUID `param:"id" validate:"required,min=1"`
}
