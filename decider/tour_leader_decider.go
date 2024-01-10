package decider

import (
	"fmt"

	"github.com/gen2brain/dlgs"
	ui "github.com/mortazavian/Hotel-Reservation-Go/UI"
	"github.com/mortazavian/Hotel-Reservation-Go/logic"
)

func TourLeaderMenu() {
	userInput, _, err := dlgs.List(ui.HotelName, "Select one item below:", []string{"Add Tour"})
	if err != nil {
		fmt.Println(err)
	}

	switch userInput {
	case "Add Tour":
		logic.AddTour()
	}
}
