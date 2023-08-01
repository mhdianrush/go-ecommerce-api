package entities

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	Id               string          `gorm:"primaryKey;not null;uniqueIndex;size:40"`
	ParentId         string          `gorm:"size:40;index"`
	UserId           string          `gorm:"size:40;index"`
	Sku              string          `gorm:"size:200;index"`
	Name             string          `gorm:"size:200"`
	Slug             string          `gorm:"size:200"`
	ShortDescription string          `gorm:"type:text"`
	Description      string          `gorm:"type:text"`
	Price            decimal.Decimal `gorm:"type:decimal(16,2)"`
	Weight           decimal.Decimal `gorm:"type:decimal(10,2)"`
	Status           int             `gorm:"default:0"`
	Stock            int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	User             User
	ProductImage     []ProductImage
	Categories       []Category `gorm:"many2many:product_categories"`
}
