package entities

import "time"

type Category struct {
	Id        string `gorm:"primaryKey;not null;uniqueIndex;size:40"`
	SectionId string `gorm:"size:40;index"`
	ParentId  string `gorm:"size:40"`
	Name      string `gorm:"size:200"`
	Slug      string `gorm:"size:200"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Section   Section
	Products  []Product `gorm:"many2many:product_categories"`
}
