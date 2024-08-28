package entity

import (
	"time"

	"github.com/google/uuid"
)

type News struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text;not null" json:"description"`
	ImageURL    string    `gorm:"type:text" json:"imageURL"`
	CreatedAt   time.Time `gorm:"type:timestamptz;not null" json:"createdAt"`
}