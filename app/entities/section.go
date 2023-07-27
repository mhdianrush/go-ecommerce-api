package entities

import "time"

type Section struct {
	Id         string `gorm:"primaryKey;not null;uniqueIndex;size:40"`
	Name       string `gorm:"size:200"`
	Slug       string `gorm:"size:200"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Categories []Category
}
