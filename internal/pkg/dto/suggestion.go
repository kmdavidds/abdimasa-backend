package dto

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type CreateSuggestionRequest struct {
	Name          string                `form:"name" validate:"required,min=1,max=255"`
	Description   string                `form:"description" validate:"required,min=1"`
	Attachment1   *multipart.FileHeader `form:"attachment1"`
}

type DeleteSuggestionRequest struct {
	ID uuid.UUID `param:"id" validate:"required,min=1"`
}
