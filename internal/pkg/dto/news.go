package dto

import "github.com/google/uuid"

type CreateNewsRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=255"`
	Description string `json:"description" validate:"required"`
	ImageURL    string `json:"imageURL" validate:"url"`
}

type GetNewsByIDRequest struct {
	ID uuid.UUID `param:"id" validate:"required"`
}

type UpdateNewsRequest struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Title       string    `json:"title" validate:"min=1,max=255"`
	Description string    `json:"description" validate:"min=1"`
	ImageURL    string    `json:"imageURL" validate:"url"`
}

type DeleteNewsRequest struct {
	ID uuid.UUID `param:"id" validate:"required"`
}