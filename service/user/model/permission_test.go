package model

import (
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestPermisson(t *testing.T) {
	db, _ := gorm.Open(mysql.Open("root:wangtao@tcp(127.0.0.1:3306)/e5Code?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"), &gorm.Config{})
	// db.AutoMigrate(&User{}, &Permission{})
	ids := []string{"9fa4d61a-0b04-464b-86e5-60101d263113"}
	permissions := []*Permission{}
	db.Where("user_id in ?", ids).First(&permissions)
	for _, v := range permissions {
		fmt.Printf("v: %v\n", v)
	}
}
