package model

import (
	"time"
)

type Shortlink struct {
	ID          string    `gorm:"primaryKey;type:varchar(20)" json:"id"`
	OriginalURL string    `gorm:"type:text;uniqueIndex;not null" json:"original_url"`
	ShortURL    string    `gorm:"-" json:"short_url,omitempty"` // Only used for response, not persisted
	CreatedAt   time.Time `json:"created_at"`
}
