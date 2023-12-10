package database

import (
	"fmt"
	"github.com/go-fiber-crud/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	dsn := "root:password@tcp(127.0.0.1:3306)/go_fiber?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	helper.PanicHandler(err)
	fmt.Println("Database connected")

}
