package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB = nil
var err error

func ConnectToDatabase() *gorm.DB {
	// Connect to database
	if db != nil && err == nil {
		return db
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=5432 sslmode=disable TimeZone=Asia/Shanghai", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	} else {
		fmt.Println("DB connection SUCCESS")
		return db
	}

}
