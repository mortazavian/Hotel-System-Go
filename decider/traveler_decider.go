package decider

import (
	"github.com/gen2brain/dlgs"
	ui "github.com/mortazavian/Hotel-Reservation-Go/UI"
	"github.com/mortazavian/Hotel-Reservation-Go/logic"
)

func TravelerDecider() {
	userInput, _, err := dlgs.List(ui.HotelName, "Select one item below:", []string{"Sign Up", "Login", "Back"})

	if userInput == "Sign Up" {
		logic.TravelerSignUp()
		TravelerDecider()
	} else if userInput == "Login" {
		logic.UserLogin()
		TravelerDecider()
	} else if userInput == "Back" {
		DecideRole()
	}

	_ = err
}
