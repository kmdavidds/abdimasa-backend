package entity

import (
	"time"

	"github.com/google/uuid"
)

type Place struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(255);not null" json:"name"`
	Location     string    `gorm:"type:varchar(255);not null" json:"location"`
	Description  string    `gorm:"type:text;not null" json:"description"`
	Address      string    `gorm:"type:varchar(255);not null" json:"address"`
	OpeningHours string    `gorm:"type:varchar(10)" json:"openingHours"`
	ClosingHours string    `gorm:"type:varchar(10)" json:"closingHours"`
	EntryPrice   string    `gorm:"type:varchar(255)" json:"entryPrice"`
	ImageURL1    string    `gorm:"type:text" json:"imageURL1"`
	ImageURL2    string    `gorm:"type:text" json:"imageURL2"`
	ImageURL3    string    `gorm:"type:text" json:"imageURL3"`
	MapURL       string    `gorm:"type:text" json:"mapURL"`
	Rating       int       `gorm:"type:int8;default:0" json:"rating"`
	CreatedAt    time.Time `gorm:"type:timestamptz;not null" json:"createdAt"`
	Reviews      []Review  `json:"reviews"`
}
