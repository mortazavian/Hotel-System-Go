package travelerlogic

import (
	"fmt"
	"time"

	"github.com/gen2brain/dlgs"
	ui "github.com/mortazavian/Hotel-Reservation-Go/UI"
	"github.com/mortazavian/Hotel-Reservation-Go/database"
	"github.com/mortazavian/Hotel-Reservation-Go/models"
)

func SignUp() {

	fmt.Println("1234343")

	makeNewTraveler()

}

func makeNewTraveler() {
	var traveler models.Traveler
	travelerInformation(&traveler)
}

func travelerInformation(*models.Traveler) {
	firstName, _, _ := dlgs.Entry(ui.HotelName, "Please enter your first name:", "")
	lastName, _, _ := dlgs.Entry(ui.HotelName, "Please enter your last name:", "")
	username, _, _ := dlgs.Entry(ui.HotelName, "Please enter your username:", "")
	password, _, _ := dlgs.Password(ui.HotelName, "Please enter your password:")
	dateOfBirth, _, _ := dlgs.Date(ui.HotelName, "Please enter your date of birth:", time.Now())
	nationalID, _, _ := dlgs.Entry(ui.HotelName, "Please enter your national id:", "")
	phoneNumber, _, _ := dlgs.Entry(ui.HotelName, "Please enter your phone number:", "")
	email, _, _ := dlgs.Entry(ui.HotelName, "Please enter your email:", "")

	traveler := models.Traveler{
		FirstName:   firstName,
		LastName:    lastName,
		Username:    username,
		Password:    password,
		DateofBirth: dateOfBirth,
		NationalID:  nationalID,
		PhoneNumber: phoneNumber,
		Email:       email,
	}

	fmt.Println(traveler)

	db := database.NewGormPostgres()

	db.Create(&traveler)

	// add it to database!!!!!!!!!!!!!!!!!!

}
