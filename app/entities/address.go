package entities

import "time"

type Address struct {
	Id         string `gorm:"primaryKey;not null;uniqueIndex;size:40"`
	UserId     string `gorm:"size:40;index"`
	FirstName  string `gorm:"size:200"`
	Address1   string `gorm:"size:200"`
	Address2   string `gorm:"size:200"`
	CityId     string `gorm:"size:200"`
	ProvinceId string `gorm:"size:200"`
	Phone      string `gorm:"size:200"`
	Email      string `gorm:"size:200"`
	PostCode   string `gorm:"size:200"`
	IsPrimary  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       User
}
