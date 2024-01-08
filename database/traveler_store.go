package database

import (
	"github.com/mortazavian/Hotel-Reservation-Go/models"
)

func InsertTraveler(traveler *models.Traveler) {

	result := Instance.Create(&traveler)
	_ = result

}

func GetUser(username, password string) (models.Traveler, error) {
	user := models.Traveler{}
	result := Instance.Where("username = ?", username).Where("password = ?", password).First(&user)
	if result.Error != nil {
		return models.Traveler{}, result.Error
	}
	return user, nil
}
