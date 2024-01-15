package logic

import (
	"fmt"

	"github.com/gen2brain/dlgs"
	ui "github.com/mortazavian/Hotel-Reservation-Go/UI"
)

func TourLeaderMenu() {
	userInput, _, err := dlgs.List(ui.HotelName, "Select one item below:", []string{"Add Tour"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(userInput)

	switch userInput {
	case "Add Tour":
		MakeNewTour()
	}
}
