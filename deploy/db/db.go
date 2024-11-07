package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDBConnection() (*gorm.DB, error) {

	if DB != nil {
		return DB, nil
	}

	dsn := "root:qwerty@tcp(127.0.0.1:3306)/vkakids?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %v", err)
	}

	fmt.Println("Success connection to DB")
	DB = db
	return DB, nil

}
