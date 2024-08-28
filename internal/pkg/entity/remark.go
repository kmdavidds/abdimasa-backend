package entity

import (
	"time"

	"github.com/google/uuid"
)

type Remark struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Occupation  string    `gorm:"type:varchar(255);not null" json:"occupation"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Gender      string    `gorm:"type:varchar(255);not null" json:"gender"`
	CreatedAt   time.Time `gorm:"type:timestamptz;not null" json:"createdAt"`
}
