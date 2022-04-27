package model

import "time"

type User struct {
	ID        string `gorm:"primaryKey;type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string `gorm:"uniqueIndex; not null; type:varchar(255)"`
	Account   string `gorm:"uniqueIndex; not null; type:varchar(255)"`
	Name      string
	Bio       string
	Password  string `gorm:"not null"`
}
