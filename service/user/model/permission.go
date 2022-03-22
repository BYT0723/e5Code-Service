package model

import "time"

type Permission struct {
	ID         string `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserID     string `gorm:"uniqueIndex:idx_userid_projectid; not null; type:varchar(255)"`
	ProjectID  string `gorm:"uniqueIndex:idx_userid_projectid; not null; type:varchar(255)"`
	Permission int
}
