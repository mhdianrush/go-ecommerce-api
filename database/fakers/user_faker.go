package fakers

import (
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"github.com/mhdianrush/go-ecommerce-api/app/entities"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *entities.User {
	return &entities.User{
		Id:            uuid.New().String(),
		FirstName:     faker.FirstName(),
		LastName:      faker.LastName(),
		Email:         faker.Email(),
		Password:      "$2a$12$1fsWqbcrWibjml0yDUXLPuLl0x99b0Iju4FQJy2Mvb670f/uCaMtK",
		RememberToken: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
	}
}
