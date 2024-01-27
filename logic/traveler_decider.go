package logic

import (
	"fmt"

	"github.com/gen2brain/dlgs"
	ui "github.com/mortazavian/Hotel-Reservation-Go/UI"
)

func TravelerDecider() {
	options := []string{}

	if loggedUser.FirstName == "" {
		options = append(options, "Sign Up", "Login", "Back")
	} else {
		options = append(options, "Reserve", "Request Cleaning", "Cancel Reservation", "Extend Reservation")
	}

	userInput, _, err := dlgs.List(ui.HotelName, "Select one item below:", options)
	if err != nil {
		fmt.Println(err)
	}

	if userInput == "Sign Up" {
		TravelerSignUp()
		TravelerDecider()
	} else if userInput == "Login" {
		UserLogin()
		TravelerDecider()
	} else if userInput == "Back" {
		DecideRole()
	} else if userInput == "Reserve" {
		MakeNewReservation()
	} else if userInput == "Request Cleaning" {
		MakeNewCleaningRequest()
	} else if userInput == "Cancel Reservation" {
		CancelReservation()
	} else if userInput == "Extend Reservation" {
		ExtendReservation()
	}

	options = nil
}
