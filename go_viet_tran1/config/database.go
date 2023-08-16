package config

import (
	"fmt"
	// "log"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

const (
	DB_USERNAME = "myuser"
	DB_PASSWORD = "mypassword"
	DB_NAME = "mydatabase"
	DB_HOST = "mysql"
	DB_PORT = "3306"
)

func DatabaseMySqlConnection() *gorm.DB {

	var err error
	dsn := DB_USERNAME +":"+ DB_PASSWORD +"@tcp"+ "(" + DB_HOST + ":" + DB_PORT +")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
 
	if err != nil {
	   fmt.Println("Error connecting to database : error=%v", err)
	   return nil
	}
 
	return db
}
