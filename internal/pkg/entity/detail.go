package entity

import "time"

type Detail struct {
	ID        uint8     `gorm:"type:uuid;primaryKey" json:"id"`
	Slug      string    `gorm:"type:varchar(255);not null" json:"slug"`
	Value     string    `gorm:"type:text;" json:"value"`
	CreatedAt time.Time `gorm:"type:timestamptz;not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamptz;not null" json:"updatedAt"`
}
