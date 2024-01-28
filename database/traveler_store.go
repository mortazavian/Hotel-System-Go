package database

import (
	"github.com/mortazavian/Hotel-Reservation-Go/models"
)

func InsertTraveler(traveler *models.Traveler) {

	result := Instance.Create(&traveler)
	_ = result

}

func GetUser(username string) (models.Traveler, error) {
	user := models.Traveler{}
	result := Instance.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return models.Traveler{}, result.Error
	}
	return user, nil
}
