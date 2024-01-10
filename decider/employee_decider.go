package decider

import (
	"fmt"

	"github.com/gen2brain/dlgs"
	ui "github.com/mortazavian/Hotel-Reservation-Go/UI"
	"github.com/mortazavian/Hotel-Reservation-Go/logic"
)

func EmployeeDecider() {
	userInput, _, err := dlgs.List(ui.HotelName, "Please select what you want.", []string{"Sign Up", "Manager",
		"Head of Facilities", "Head of Reservation", "Head of Restaurant", "Head of Pool", "Tour Leader",
		"Head of Cleaning", "Cleaner", "EXIT"})

	if err != nil {
		fmt.Println(err)
	}

	switch userInput {
	case "Sign Up":
		logic.EmployeeSignUp()
	case "Manager":
		// TODO
		return
	case "Head of Facilities":
		// TODO
		return
	case "Head of Reservation":
		// TODO
		return
	case "Head of Restaurant":
		// TODO
		return
	case "Head of Pool":
		// TODO
		return
	case "Tour Leader":
		TourLeaderMenu()
		return
	case "Head of Cleaning":
		// TODO
		return
	case "Cleaner":
		// TODO
		return
	case "EXIT":
		//TODO

	}
}
