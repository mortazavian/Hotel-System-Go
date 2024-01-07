package decider

import (
	"github.com/gen2brain/dlgs"
	ui "github.com/mortazavian/Hotel-Reservation-Go/UI"
)

func DecideRole() {
	userInput, _, err := dlgs.List(ui.HotelName, "Select your role", []string{"Traveler", "Manager",
		"Head of Facilities", "Head of Reservation", "Head of Restaurant", "Head of Pool", "Tour Leader",
		"Head of Cleaning", "Cleaner", "EXIT"})
	if err != nil {
		panic(err)
	}
	switch userInput {
	case "Traveler":
		TravelerDecider()
	case "EXIT":
		return
	}

}
