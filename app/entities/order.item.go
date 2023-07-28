package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderItem struct {
	Id              string          `gorm:"primaryKey;not null;uniqueIndex;size:40"`
	OrderId         string          `gorm:"size:40;index"`
	ProductId       string          `gorm:"size:40;index"`
	BasePrice       decimal.Decimal `gorm:"type:decimal(16,2)"`
	BaseTotal       decimal.Decimal `gorm:"type:decimal(16,2)"`
	TaxAmount       decimal.Decimal `gorm:"type:decimal(16,2)"`
	TaxPercent      decimal.Decimal `gorm:"type:decimal(10,2)"`
	DiscountAmount  decimal.Decimal `gorm:"type:decimal(16,2)"`
	DiscountPercent decimal.Decimal `gorm:"type:decimal(10,2)"`
	SubTotal        decimal.Decimal `gorm:"type:decimal(16,2)"`
	Sku             string          `gorm:"size:40;index"`
	Name            string          `gorm:"size:300"`
	Weight          decimal.Decimal `gorm:"type:decimal(10,2)"`
	Quantity        int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Order           Order
	Product         Product
}
