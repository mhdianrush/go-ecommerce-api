package fakers

import (
	"math"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/mhdianrush/go-ecommerce-api/app/entities"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var logger = logrus.New()

func ProductFaker(db *gorm.DB) *entities.Product {
	user := UserFaker(db)
	if err := db.Create(&user).Error; err != nil {
		logger.Printf("failed create user data %s", err.Error())
	}
	return &entities.Product{
		Id:               uuid.New().String(),
		UserId:           user.Id,
		Sku:              slug.Make(faker.Name()),
		Name:             faker.Name(),
		Slug:             slug.Make(faker.Name()),
		ShortDescription: faker.Paragraph(),
		Description:      faker.Paragraph(),
		Price:            decimal.NewFromFloat(fakePrice()),
		Weight:           decimal.NewFromFloat(rand.Float64()),
		Status:           1,
		Stock:            rand.Intn(200),
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		DeletedAt:        gorm.DeletedAt{},
	}
}

func fakePrice() float64 {
	return Precision(rand.Float64()*math.Pow10(rand.Intn(8)), rand.Intn(2)+1)
}

func Precision(val float64, pre int) float64 {
	div := math.Pow10(pre)
	return float64(int64(val*div)) / div
}
