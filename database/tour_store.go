package database

import "github.com/mortazavian/Hotel-Reservation-Go/models"

func InsertTour(tour *models.Tour) {
	result := Instance.Create(&tour)
	_ = result
}
