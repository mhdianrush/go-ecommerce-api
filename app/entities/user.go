package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id            string `gorm:"primaryKey;not null;uniqueIndex;size:40"`
	FirstName     string `gorm:"not null;size:200"`
	LastName      string `gorm:"not null;size:200"`
	Email         string `gorm:"not null;uniqueIndex;size:200"`
	Password      string `gorm:"not null;size:200"`
	RememberToken string `gorm:"not null;size:200"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	Addresses     []Address
}
