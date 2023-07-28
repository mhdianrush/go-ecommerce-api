package entities

import "time"

type OrderCustomer struct {
	Id         string `gorm:"primaryKey;not null;uniqueIndex;size:40"`
	UserId     string `gorm:"size:40;index"`
	OrderId    string `gorm:"size:40;index"`
	FirstName  string `gorm:"not null;size:200"`
	LastName   string `gorm:"not null;size:200"`
	CityId     string `gorm:"size:200"`
	ProvinceId string `gorm:"size:200"`
	Address1   string `gorm:"size:200"`
	Address2   string `gorm:"size:200"`
	Phone      string `gorm:"size:200"`
	Email      string `gorm:"size:200"`
	PostCode   string `gorm:"size:200"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       User
	Order      Order
}
