package database

import (
	"github.com/mortazavian/Hotel-Reservation-Go/models"
)

func InsertTour(tour *models.Tour) {
	result := Instance.Create(&tour)
	_ = result
}

func GetTours() []string {
	tours := []models.Tour{}
	result := Instance.Find(&tours)
	_ = result

	strArrToReturn := []string{}

	for _, tour := range tours {
		strArrToReturn = append(strArrToReturn, tour.TourName)
	}

	return strArrToReturn
}

func DeleteTour(tourName string) {
	tour := models.Tour{TourName: tourName}
	result := Instance.Where("tour_name = ?", tourName).Delete(&tour)
	_ = result
}
