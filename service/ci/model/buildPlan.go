package model

import "time"

type BuildPlan struct {
	ID         string `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string `gorm:"uniqueIndex: idx_name_projectid; not null; type: varchar(255)"`
	ProjectID  string `gorm:"uniqueIndex: idx_name_projectid; not null; type: varchar(255)"`
	Context    string
	Dockerfile string
	Version    string
	Tag        string
	Already    bool
}
