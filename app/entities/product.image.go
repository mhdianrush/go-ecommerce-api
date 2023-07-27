package entities

import "time"

type ProductImage struct {
	Id         string `gorm:"primaryKey;not null;uniqueIndex;size:40"`
	ProductId  string `gorm:"size:40;index"`
	Path       string `gorm:"type:text"`
	ExtraLarge string `gorm:"type:text"`
	Large      string `gorm:"type:text"`
	Medium     string `gorm:"type:text"`
	Small      string `gorm:"type:text"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Product    Product
}
