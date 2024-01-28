package database

import "github.com/mortazavian/Hotel-Reservation-Go/models"

func InsertRoom(room *models.Room) {
	result := Instance.Create(&room)
	_ = result
}
