package database

import (
	"strings"

	"github.com/mortazavian/Hotel-Reservation-Go/models"
)

func GetAllUserReservations(traveler models.Traveler) []models.Reservation {
	reservations := []models.Reservation{}
	Instance.Where("user_id = ?", traveler.ID).Find(&reservations)
	return reservations
}

func InsertCleaningRequest(reserveTime string, traveler models.Traveler) {
	reserveTime = strings.TrimRight(reserveTime, "+330")

	var reservation models.Reservation

	Instance.Where("date = ?", reserveTime).First(&reservation)

	var cleaningRequest models.CleaningRequest
	cleaningRequest.UserID = traveler.ID
	cleaningRequest.ReservationID = reservation.ID

	result := Instance.Create(&cleaningRequest)
	_ = result
}
