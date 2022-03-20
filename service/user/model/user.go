package model

import "time"

type User struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string `gorm:"uniqueIndex; not null; type:varchar(255)"`
	Accout    string `gorm:"uniqueIndex; not null; type:varchar(255)"`
	Name      string
	Password  string `gorm:"not null"`
}
