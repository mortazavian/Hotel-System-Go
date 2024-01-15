package database

import "github.com/mortazavian/Hotel-Reservation-Go/models"

func InsertFood(food *models.Food) {
	result := Instance.Create(&food)
	_ = result
}
