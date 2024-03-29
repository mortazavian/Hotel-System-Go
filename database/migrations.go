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

	err = db.AutoMigrate(&models.Tour{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&models.Employee{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&models.Food{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&models.Room{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&models.Reservation{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&models.CleaningRequest{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&models.PoolSans{})
	if err != nil {
		return
	}

}
