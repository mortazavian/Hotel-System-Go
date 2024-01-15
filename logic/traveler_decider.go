package logic

import (
	"github.com/gen2brain/dlgs"
	ui "github.com/mortazavian/Hotel-Reservation-Go/UI"
)

func TravelerDecider() {
	userInput, _, err := dlgs.List(ui.HotelName, "Select one item below:", []string{"Sign Up", "Login", "Back"})

	if userInput == "Sign Up" {
		TravelerSignUp()
		TravelerDecider()
	} else if userInput == "Login" {
		UserLogin()
		TravelerDecider()
	} else if userInput == "Back" {
		DecideRole()
	}

	_ = err
}
