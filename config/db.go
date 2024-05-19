package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DatabaseInit() {
	host := "localhost"
	user := "root"
	password := "pisangkukus"
	dbName := "go_echo_api"
	port := 3306

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	database, e = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}
	fmt.Println("Successfully connected to MySQL database")
}

func DB() *gorm.DB {
	return database
}
