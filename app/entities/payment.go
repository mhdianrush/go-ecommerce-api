package entities

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Payment struct {
	Id                   string          `gorm:"primaryKey;not null;uniqueIndex;size:40"`
	OrderId              string          `gorm:"size:40;index"`
	Number               string          `gorm:"size:200;index"`
	Method               string          `gorm:"size:200"`
	Status               string          `gorm:"size:200"`
	Token                string          `gorm:"size:200;index"`
	Payload              string          `gorm:"type:text"`
	PaymentType          string          `gorm:"size:200"`
	VirtualAccountNumber string          `gorm:"size:200"`
	BillCode             string          `gorm:"size:200"`
	BillKey              string          `gorm:"size:200"`
	Amount               decimal.Decimal `gorm:"type:decimal(16,2)"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt
	Order                Order
}
