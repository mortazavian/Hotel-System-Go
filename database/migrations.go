package database

import (
	"github.com/mortazavian/Hotel-Reservation-Go/models"
	"gorm.io/gorm"
)

func MakeMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.Traveler{})
	if err != nil {
		return
	}

}
