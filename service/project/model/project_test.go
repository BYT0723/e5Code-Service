package model

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestProject(t *testing.T) {
	db, _ := gorm.Open(mysql.Open("root:wangtao@tcp(127.0.0.1:3306)/e5Code?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"), &gorm.Config{})
	db.AutoMigrate(&Project{})
}
