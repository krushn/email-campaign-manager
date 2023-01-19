package migrations

import (
	"email-campaign/db"
	"email-campaign/models"
	"log"
)

func Migrate() {

	err := db.GetDB().AutoMigrate(&models.User{}, &models.Campaign{}, &models.Subscribers{})

	if err != nil {
		log.Panic(err)
	}
}
