package models

import (
	"gorm.io/gorm"
)

// User Model
type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FullName  string `json:"full_name"`
	Email     string `json:"email" gorm:"unique"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}
