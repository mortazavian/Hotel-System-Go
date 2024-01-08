package logic

import (
	"fmt"
	"log"
	"time"

	"github.com/gen2brain/dlgs"
	ui "github.com/mortazavian/Hotel-Reservation-Go/UI"
	"github.com/mortazavian/Hotel-Reservation-Go/database"
	"github.com/mortazavian/Hotel-Reservation-Go/models"
)

func SignUp() {
	MakeNewUser()
}

func MakeNewUser() {
	user := models.Traveler{}
	UserInformation(&user)
}

func UserInformation(traveler *models.Traveler) {
	firstName, _, err := dlgs.Entry(ui.HotelName, "Enter you first name:", "")
	if err != nil {
		log.Fatal(err)
	}

	lastName, _, err := dlgs.Entry(ui.HotelName, "Enter you last name:", "")
	if err != nil {
		log.Fatal(err)
	}

	username, _, err := dlgs.Entry(ui.HotelName, "Enter you username:", "")
	if err != nil {
		log.Fatal(err)
	}

	password, _, err := dlgs.Password(ui.HotelName, "Enter you password:")
	if err != nil {
		log.Fatal(err)
	}

	dataOfBirth, _, err := dlgs.Date(ui.HotelName, "Enter the date of birth:", time.Now())
	if err != nil {
		log.Fatal(err)
	}

	nationalID, _, err := dlgs.Entry(ui.HotelName, "Enter your national id:", "")
	if err != nil {
		log.Fatal(err)
	}

	phoneNumber, _, err := dlgs.Entry(ui.HotelName, "Enter your phone number: (09171234567)", "")
	if err != nil {
		log.Fatal(err)
	}

	email, _, err := dlgs.Entry(ui.HotelName, "Enter your email: (x@y.com)", "")
	if err != nil {
		log.Fatal(err)
	}

	traveler.FirstName = firstName
	traveler.LastName = lastName
	traveler.Username = username
	traveler.Password = password
	traveler.DateofBirth = dataOfBirth
	traveler.NationalID = nationalID
	traveler.PhoneNumber = phoneNumber
	traveler.Email = email

	database.InsertTraveler(traveler)
}

func UserLogin() {
	username, _, err := dlgs.Entry(ui.HotelName, "enter your username: ", "")
	if err != nil {
		log.Fatal(err)
	}
	password, _, err := dlgs.Password(ui.HotelName, "enter your password: ")
	if err != nil {
		log.Fatal(err)
	}
	user, err := database.GetUser(username, password)
	if err != nil {
		fmt.Println("--------------------")
		fmt.Println(err)
		fmt.Println("--------------------")
	}
	if err == nil {
		fmt.Println(user)
		// decider.TravelerDecider()
	}
}
