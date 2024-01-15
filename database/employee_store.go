package database

import "github.com/mortazavian/Hotel-Reservation-Go/models"

func InsertEmployee(employee *models.Employee) {
	result := Instance.Create(&employee)
	_ = result
}

func GetEmployee(username string) (models.Employee, error) {
	employee := models.Employee{}
	result := Instance.Where("username = ?", username).First(&employee)
	if result.Error != nil {
		return models.Employee{}, result.Error
	}
	return employee, nil
}
 