package database

import "github.com/mortazavian/Hotel-Reservation-Go/models"

func InsertEmployee(employee *models.Employee) {
	result := Instance.Create(&employee)
	_ = result
}
