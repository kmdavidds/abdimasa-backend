package entity

import (
	"time"

	"github.com/google/uuid"
)

type Suggestion struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name          string    `gorm:"type:varchar(255);not null" json:"name"`
	Description   string    `gorm:"type:text;not null" json:"description"`
	AttachmentURL string    `gorm:"type:text" json:"attachmentURL"`
	CreatedAt     time.Time `gorm:"type:timestamptz;not null" json:"createdAt"`
}