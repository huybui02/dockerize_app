package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

const (
	host     = "mysql"
	port     = 3306
	user     = "myuser"
	password = "mypassword"
	dbName   = "mydatabase"
)


func DatabaseMySqlConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
