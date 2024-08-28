package dto

import "github.com/google/uuid"

type CreateSuggestionRequest struct {
	Name          string `json:"name" validate:"required,min=1,max=255"`
	Description   string `json:"description" validate:"required,min=1"`
	AttachmentURL string `json:"attachmentURL" validate:"min=1"`
}

type DeleteSuggestionRequest struct {
	ID uuid.UUID `param:"id" validate:"required,min=1"`
}
