package model

import (
	userModel "e5Code-Service/service/user/model"
	"time"
)

type Image struct {
	ID          string `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	ImageID     string
	ProjectID   string
	BuildPlanID string
	BuilderID   string
	Builder     *userModel.User `gorm:"migration"`
}
