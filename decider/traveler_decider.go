package decider

import (
	"github.com/gen2brain/dlgs"
	ui "github.com/mortazavian/Hotel-Reservation-Go/UI"
	"github.com/mortazavian/Hotel-Reservation-Go/logic/travelerlogic"
)

func TravelerDecider() {
	userInput, _, err := dlgs.List(ui.HotelName, "Select one item below:", []string{"Sign Up", "b", "Back"})

	if userInput == "Sign Up" {
		travelerlogic.SignUp()
	}

	_ = err
}
