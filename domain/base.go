package domain

import (
	"time"

	guuid "github.com/google/uuid"
)

//Base means the basic of a domain
type Base struct {
	ID        guuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NewBase creates a new Base
func NewBase() Base {
	return Base{
		ID:        guuid.New(),
		CreatedAt: time.Now(),
	}
}
