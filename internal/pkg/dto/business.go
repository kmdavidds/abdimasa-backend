package dto

import "github.com/google/uuid"

type CreateBusinessRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=255"`
	Location    string `json:"location" validate:"required,min=1,max=255"`
	Description string `json:"description" validate:"required,min=1"`
	Address     string `json:"address" validate:"required,min=1,max=255"`
	PriceRange  string `json:"priceRange" validate:"min=1,max=255"`
	ImageURL1   string `json:"imageURL1" validate:"url"`
	ImageURL2   string `json:"imageURL2" validate:"url"`
	ImageURL3   string `json:"imageURL3" validate:"url"`
	Contact     string `json:"contact" validate:"min=1,max=255"`
	MapURL      string `json:"mapURL" validate:"url"`
	Rating      int    `json:"rating" validate:"min=0,max=50"`
}

type UpdateBusinessRequest struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Name        string    `json:"name" validate:"min=1,max=255"`
	Location    string    `json:"location" validate:"min=1,max=255"`
	Description string    `json:"description" validate:"min=1"`
	Address     string    `json:"address" validate:"min=1,max=255"`
	PriceRange  string    `json:"priceRange" validate:"min=1,max=255"`
	ImageURL1   string    `json:"imageURL1" validate:"url"`
	ImageURL2   string    `json:"imageURL2" validate:"url"`
	ImageURL3   string    `json:"imageURL3" validate:"url"`
	Contact     string    `json:"contact" validate:"min=1,max=255"`
	MapURL      string    `json:"mapURL" validate:"url"`
	Rating      int       `json:"rating" validate:"min=0,max=50"`
}

type DeleteBusinessRequest struct {
	ID uuid.UUID `param:"id" validate:"required"`
}