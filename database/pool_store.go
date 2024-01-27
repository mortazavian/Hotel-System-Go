package database

import "github.com/mortazavian/Hotel-Reservation-Go/models"

func InsertPoolSans(sans *models.PoolSans) {
	Instance.Create(sans)
}
