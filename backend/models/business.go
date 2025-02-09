package models

import (
	"time"
)

type Business struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	Email      string    `gorm:"uniqueIndex" json:"email"`
	Password   string    `json:"-"`
	Address    string    `json:"address"`
	PostalCode string    `json:"postal_code"`
	Logo       string    `json:"logo"` // Store file path
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
