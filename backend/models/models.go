package models

import "gorm.io/gorm"

// User model
// type User struct {
// 	gorm.Model
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// 	Email     string `json:"email" gorm:"unique"`
// 	Password  string `json:"-"`
// 	Role      string `json:"role"` // "customer" or "business"
// }

// // Business model
// type Business struct {
// 	gorm.Model
// 	Name        string `json:"name"`
// 	Address     string `json:"address"`
// 	PostalCode  string `json:"postal_code"`
// 	OwnerID     uint   `json:"owner_id"` // Foreign Key: UserID
// 	Phone       string `json:"phone"`
// 	Description string `json:"description"`
// 	Photos      string `json:"photos"` // Store photo URLs (comma-separated)
// }

// Appointment model
type Appointment struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	BusinessID  uint   `json:"business_id"`
	DateTime    string `json:"date_time"`
	Status      string `json:"status"`
	Description string `json:"description"`
}
