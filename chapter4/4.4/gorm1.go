package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	Id       int    `json:"id"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "user"
}

func main() {
	dns := "root:password@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Init mysql failed , err:", err)
		return
	}
	GormUser := User{
		Name:     "jaluik",
		Phone:    "18580001234",
		Password: "pwd",
	}
	db.Create(&GormUser)
	var GormQueryUser = new(User)
	db.Where("id = ?", 1).Find(&GormQueryUser)
	fmt.Println(GormQueryUser)
}
