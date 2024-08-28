package dto

import "github.com/google/uuid"

type CreateActivityRequest struct {
	Title    string `json:"title" validate:"required,min=1,max=255"`
	ImageURL string `json:"imageURL" validate:"required,url"`
	Date     string `json:"date" validate:"required,len=10"`
	Time     string `json:"time" validate:"min=9,max=10"`
	Location string `json:"location" validate:"required,min=1,max=255"`
}

type UpdateActivityRequest struct {
	ID       uuid.UUID `json:"id" validate:"min=1"`
	Title    string    `json:"title" validate:"min=1,max=255"`
	ImageURL string    `json:"imageURL" validate:"url"`
	Date     string    `json:"date" validate:"len=10"`
	Time     string    `json:"time" validate:"min=9,max=10"`
	Location string    `json:"location" validate:"min=1,max=255"`
}

type DeleteActivityRequest struct {
	ID uuid.UUID `param:"id" validate:"required,min=1"`
}
