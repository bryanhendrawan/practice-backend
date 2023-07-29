package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	dsn := "user:password@tcp(127.0.0.1:3306)/practice_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed connect to database MySQL")
	}

	fmt.Println("Connected to MYSQL Database")

	DB = db
}
