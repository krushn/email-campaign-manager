package db

import (
	"email-campaign/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func ConnectDb() *gorm.DB {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error: failed to connect with database")
	}

	db.AutoMigrate(&models.Campaign{}, &models.User{}, &models.Subscribers{})

	return db
}
