package entity

import "time"

type BookGorm struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `gorm:"not null;type:varchar(75)"`
	Author    string `gorm:"not null;type:varchar(50)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
