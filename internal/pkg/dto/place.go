package dto

import "github.com/google/uuid"

type CreatePlaceRequest struct {
	Name         string `json:"name" validate:"required,min=1,max=255"`
	Location     string `json:"location" validate:"required,min=1,max=255"`
	Description  string `json:"description" validate:"required,min=1"`
	Address      string `json:"address" validate:"required,min=1,max=255"`
	OpeningHours string `json:"openingHours" validate:"min=9,max=10"`
	ClosingHours string `json:"closingHours" validate:"min=9,max=10"`
	EntryPrice   string `json:"entryPrice" validate:"min=1,max=255"`
	ImageURL1    string `json:"imageURL1" validate:"url"`
	ImageURL2    string `json:"imageURL2" validate:"url"`
	ImageURL3    string `json:"imageURL3" validate:"url"`
	MapURL       string `json:"mapURL" validate:"url"`
	Rating       int    `json:"rating" validate:"min=0,max=50"`
}

type UpdatePlaceRequest struct {
	ID           uuid.UUID `json:"id" validate:"required"`
	Name         string    `json:"name" validate:"min=1,max=255"`
	Location     string    `json:"location" validate:"min=1,max=255"`
	Description  string    `json:"description" validate:"min=1"`
	Address      string    `json:"address" validate:"min=1,max=255"`
	OpeningHours string    `json:"openingHours" validate:"min=9,max=10"`
	ClosingHours string    `json:"closingHours" validate:"min=9,max=10"`
	EntryPrice   string    `json:"entryPrice" validate:"min=1,max=255"`
	ImageURL1    string    `json:"imageURL1" validate:"url"`
	ImageURL2    string    `json:"imageURL2" validate:"url"`
	ImageURL3    string    `json:"imageURL3" validate:"url"`
	MapURL       string    `json:"mapURL" validate:"url"`
	Rating       int       `json:"rating" validate:"min=0,max=50"`
}

type DeletePlaceRequest struct {
	ID uuid.UUID `param:"id" validate:"required"`
}
