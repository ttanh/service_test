package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

var db *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	var err error
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err = gorm.Open("mysql", connectString)
	db.LogMode(os.Getenv("DB_LOG_MODE") == "1")
	return db, err
}

func InitDB() {
	db.AutoMigrate(&Token{})
	db.AutoMigrate(&Transaction{})
}
