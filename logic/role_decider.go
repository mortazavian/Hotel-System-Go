package logic

import (
	"github.com/gen2brain/dlgs"
	ui "github.com/mortazavian/Hotel-Reservation-Go/UI"
)

func DecideRole() {
	userInput, _, err := dlgs.List(ui.HotelName, "Select your role", []string{"Traveler", "Employee", "EXIT"})
	if err != nil {
		panic(err)
	}
	switch userInput {
	case "Traveler":
		TravelerDecider()
	case "Employee":
		EmployeeDecider()
	case "EXIT":
		return
	}

}
