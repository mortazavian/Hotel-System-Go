package database

import "github.com/mortazavian/Hotel-Reservation-Go/models"

func InsertFood(food *models.Food) {
	result := Instance.Create(&food)
	_ = result
}

func GetFoods() []string {
	foods := []models.Food{}
	result := Instance.Find(&foods)
	_ = result

	strArrToReturn := []string{}

	for _, food := range foods {
		strArrToReturn = append(strArrToReturn, food.Name)
	}

	return strArrToReturn
}

func DeleteFood(foodName string) {
	food := models.Food{Name: foodName}
	result := Instance.Where("name = ?", foodName).Delete(&food)
	_ = result
}
