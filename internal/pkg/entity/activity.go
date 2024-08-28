package entity

import (
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title"`
	ImageURL  string    `gorm:"type:text;not null" json:"imageURL"`
	Date      string    `gorm:"type:varchar(10);not null" json:"date"`
	Time      string    `gorm:"type:varchar(10)" json:"time"`
	Location  string    `gorm:"type:varchar(255);not null" json:"location"`
	CreatedAt time.Time `gorm:"type:timestamptz;not null" json:"createdAt"`
}