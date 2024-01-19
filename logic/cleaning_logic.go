package logic

import (
	"fmt"

	"github.com/gen2brain/dlgs"
	ui "github.com/mortazavian/Hotel-Reservation-Go/UI"
	"github.com/mortazavian/Hotel-Reservation-Go/database"
)

func MakeNewCleaningRequest() {
	allReservations := database.GetAllUserReservations(loggedUser)

	options := []string{}

	for _, reservation := range allReservations {
		options = append(options, reservation.Date.String())
	}

	userInput, _, err := dlgs.List(ui.HotelName, "Please enter which reservation you want clean for: :", options)
	if err != nil {
		fmt.Println(err)
	}

	database.InsertCleaningRequest(userInput, loggedUser)

	_, err = dlgs.MessageBox(ui.HotelName, "Click OK to continue")
	if err != nil {
		fmt.Println(err)
	}
}
