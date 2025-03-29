package database

import (
	"log"

	"github.com/killerrekt/quantstrategix/internal/model"
)

func RunMigration() {
	err := DB.AutoMigrate(model.User{})
	if err != nil {
		log.Fatalln("Failed to create User table")
	}
	err = DB.AutoMigrate(model.Subject{})
	if err != nil {
		log.Fatalln("Failed to create Subject table")
	}
	err = DB.AutoMigrate(model.Tutor{})
	if err != nil {
		log.Fatalln("Failed to create Tutor table")
	}
	err = DB.AutoMigrate(model.Booking{})
	if err != nil {
		log.Fatalln("Failed to create Booking table")
	}

	log.Println("Migration done Successfully 🚀")
}
