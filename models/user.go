package models

import (
	"time"
)

// User struct defines the model for the User entity
type User struct {
	ID        uint      `gorm:"primary_key"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Phone     string    `gorm:"unique;not null"`
	Code      string     `gorm:"unique;not null"`
	Role      string    `gorm:"not null"` // Role: 'customer' or 'admin'
	CreatedAt time.Time
	UpdatedAt time.Time
}



