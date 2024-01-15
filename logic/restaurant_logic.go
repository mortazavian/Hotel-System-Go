package logic

import (
	"fmt"

	"github.com/gen2brain/dlgs"
	ui "github.com/mortazavian/Hotel-Reservation-Go/UI"
	"github.com/mortazavian/Hotel-Reservation-Go/database"
	"github.com/mortazavian/Hotel-Reservation-Go/models"
)

func MakeNewFood() {
	food := models.Food{}
	FoodInformation(&food)
}

func FoodInformation(food *models.Food) {
	name, _, err := dlgs.Entry(ui.HotelName, "Enter the name of the food:", "")
	if err != nil {
		fmt.Println(err)
	}

	ingredients, _, err := dlgs.Entry(ui.HotelName, "Enter the ingredients in the food:", "")
	if err != nil {
		fmt.Println(err)
	}

	detail, _, err := dlgs.Entry(ui.HotelName, "Enter the detail of the food:", "")
	if err != nil {
		fmt.Println(err)
	}

	food.Name = name
	food.Ingredient = ingredients
	food.Detail = detail
	food.RestaurantHeadID = loggedEmployee.ID

	database.InsertFood(food)

	EmployeeDecider()
}
