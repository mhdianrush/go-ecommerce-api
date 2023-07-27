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
	Sku              string          `gorm:"size:150;index"`
	Name             string          `gorm:"size:200"`
	Slug             string          `gorm:"size:200"`
	Price            decimal.Decimal `gorm:"type:decimal(16,2)"`
	Stock            int
	Weight           decimal.Decimal `gorm:"type:decimal(10,2)"`
	ShortDescription string          `gorm:"size:200"`
	Description      string          `gorm:"type:text"`
	Status           int             `gorm:"default:0"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	User             User
	ProductImage     []ProductImage
}
