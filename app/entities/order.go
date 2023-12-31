package entities

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Order struct {
	Id                  string          `gorm:"primaryKey;not null;uniqueIndex;size:40"`
	UserId              string          `gorm:"size:40;index"`
	Code                string          `gorm:"size:200;index"`
	PaymentStatus       string          `gorm:"size:200;index"`
	PaymentToken        string          `gorm:"size:200;index"`
	BaseTotalPrice      decimal.Decimal `gorm:"type:decimal(16,2)"`
	TaxAmount           decimal.Decimal `gorm:"type:decimal(16,2)"`
	TaxPercent          decimal.Decimal `gorm:"type:decimal(10,2)"`
	DiscountAmount      decimal.Decimal `gorm:"type:decimal(16,2)"`
	DiscountPercent     decimal.Decimal `gorm:"type:decimal(10,2)"`
	ShippingCost        decimal.Decimal `gorm:"type:decimal(16,2)"`
	GrandTotal          decimal.Decimal `gorm:"type:decimal(16,2)"`
	Note                string          `gorm:"type:text"`
	ShippingCourier     string          `gorm:"size:200"`
	ShippingServiceName string          `gorm:"size:200"`
	ApprovedBy          string          `gorm:"size:40"`
	CancelledBy         string          `gorm:"size:40"`
	CancellationNote    string          `gorm:"size:200"`
	Status              int
	OrderDate           time.Time
	PaymentDue          time.Time
	ApprovedAt          time.Time
	CancelledAt         time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
	User                User
	OrderItems          []OrderItem
	OrderCustomer       *OrderCustomer
}
