package entity

import (
	"github.com/google/uuid"
)

type Review struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PlaceID     uuid.UUID `json:"placeID"`
	BusinessID  uuid.UUID `json:"businessID"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Title       string    `gorm:"type:varchar(255)" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Score       string    `gorm:"type:int;not null" json:"score"`
}
