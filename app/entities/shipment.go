package entities

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Shipment struct {
	Id            string          `gorm:"primaryKey;not null;uniqueIndex;size:40"`
	UserId        string          `gorm:"size:40;index"`
	OrderId       string          `gorm:"size:40;index"`
	TrackNumber   string          `gorm:"size:200;index"`
	Status        string          `gorm:"size:40;index"`
	Firstname     string          `gorm:"size:200;not null"`
	LastName      string          `gorm:"size:200;not null"`
	CityId        string          `gorm:"size:200"`
	ProvinceId    string          `gorm:"size:200"`
	Address1      string          `gorm:"size:200"`
	Address2      string          `gorm:"size:200"`
	Phone         string          `gorm:"size:200"`
	Email         string          `gorm:"size:200"`
	PostCode      string          `gorm:"size:200"`
	ShippedBy     string          `gorm:"size:40"`
	TotalWeight   decimal.Decimal `gorm:"type:decimal(10,2)"`
	TotalQuantity int
	ShippedAt     time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	User          User
	Order         Order
}
