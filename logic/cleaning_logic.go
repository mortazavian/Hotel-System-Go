package logic

import (
	"fmt"
	"strconv"

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

func AssignCleaning() {
	SeeAllCleaningRequests()
}

func SeeAllCleaningRequests() {
	cleaningRequests := database.GetAllCleaningRequests()

	cleaningRequestOptions := []string{}

	for _, cleaningRequest := range cleaningRequests {
		cleaningRequestOptions = append(cleaningRequestOptions, strconv.Itoa(int(cleaningRequest.ID)))
	}
	cleaningRequestID, _, err := dlgs.List(ui.HotelName, "Enter which one you want to assign:", cleaningRequestOptions)
	if err != nil {
		fmt.Println(err)
	}

	workers := database.GetAllWorkers()

	workersOptions := []string{}

	for _, worker := range workers {
		workersOptions = append(workersOptions, strconv.Itoa(int(worker.ID)))
	}

	workerID, _, err := dlgs.List(ui.HotelName, "Enter which worker you want to assign work to:", workersOptions)
	if err != nil {
		fmt.Println(err)
	}

	database.AssignCleaningToWorker(workerID, cleaningRequestID)

	EmployeeDecider()
}
