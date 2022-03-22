package model

import (
	"e5Code-Service/service/user/model"
	"time"
)

type Project struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"uniqueIndex:idx_name_ownerid; not null; type:varchar(255)"`
	Desc      string
	Url       string
	OwnerId   string      `gorm:"uniqueIndex:idx_name_ownerid; not null; type:varchar(255)"`
	Owner     *model.User `gorm:"migration"`
}
