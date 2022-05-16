package model

import "time"

type SSHKey struct {
  ID        string `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  Name      string `gorm:"uniqueIndex:idx_name_ownerid; not null; type:varchar(255)"`
  Key       string
  OwnerID   string `gorm:"uniqueIndex:idx_name_ownerid; not null; type:varchar(255)"`
}
