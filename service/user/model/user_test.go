package model

import (
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestUser(t *testing.T) {
	db, _ := gorm.Open(mysql.Open("root:wangtao@tcp(127.0.0.1:3306)/e5Code?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"), &gorm.Config{})
	u := &User{}
	email := "twang9739@163.com"

	if err := db.Where("email = ?", email).First(&u).Error; err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("u: %v\n", u)

	u.Name = "Wang"
	if err := db.Save(u).Error; err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("u: %v\n", u)

}
