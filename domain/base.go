package domain

import (
	"getmail/util"
	"time"
)

//Base means the basic of a domain
type Base struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
}

//NewBase creates a new Base
func NewBase() Base {
	return Base{
		ID:        util.NewID(),
		CreatedAt: time.Now(),
	}
}
