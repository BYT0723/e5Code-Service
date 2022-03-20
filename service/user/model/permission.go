package model

import "time"

type Permission struct {
	ID         string `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserID     string `gorm:"index"`
	ProjectID  string `gorm:"index"`
	Permission int
}
