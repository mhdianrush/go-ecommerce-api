package seeders

import (
	"github.com/mhdianrush/go-ecommerce-api/database/fakers"
	"gorm.io/gorm"
)

type Seeder struct {
	Seeder any
}

func RegisterSeeders(db *gorm.DB) []Seeder {
	return []Seeder{
		{
			Seeder: fakers.UserFaker(db),
		},
		{
			Seeder: fakers.ProductFaker(db),
		},
	}
}

func DatabaseSeed(db *gorm.DB) error {
	for _, seeder := range RegisterSeeders(db) {
		if err := db.Debug().Create(seeder.Seeder).Error; err != nil {
			return err
		}
	}
	return nil
}
