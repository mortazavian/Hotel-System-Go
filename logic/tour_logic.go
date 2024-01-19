package logic

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gen2brain/dlgs"
	ui "github.com/mortazavian/Hotel-Reservation-Go/UI"
	"github.com/mortazavian/Hotel-Reservation-Go/database"
	"github.com/mortazavian/Hotel-Reservation-Go/models"
)

func MakeNewTour() {
	tour := models.Tour{}
	TourInformation(&tour)
	EmployeeDecider()
}

func TourInformation(tour *models.Tour) {

	tourName, _, err := dlgs.Entry(ui.HotelName, "Enter the name of the tour:", "")
	if err != nil {
		fmt.Println(err)
	}

	StartLocation, _, err := dlgs.Entry(ui.HotelName, "Enter the start location of the tour:", "")
	if err != nil {
		fmt.Println(err)
	}

	StartLatitude, _, err := dlgs.Entry(ui.HotelName, "Enter the start latitude of the tour:", "")
	if err != nil {
		fmt.Println(err)
	}
	StartLatitudeFloat, err := strconv.ParseFloat(StartLatitude, 32)
	if err != nil {
		fmt.Println(err)
	}

	StartLongitude, _, err := dlgs.Entry(ui.HotelName, "Enter the start longitude of the tour:", "")
	if err != nil {
		fmt.Println(err)
	}
	StartLongitudeFloat, err := strconv.ParseFloat(StartLongitude, 32)
	if err != nil {
		fmt.Println(err)
	}

	EndLatitude, _, err := dlgs.Entry(ui.HotelName, "Enter the end latitude of the tour:", "")
	if err != nil {
		fmt.Println(err)
	}

	EndLatitudeFloat, err := strconv.ParseFloat(EndLatitude, 32)
	if err != nil {
		fmt.Println(err)
	}

	EndLongitude, _, err := dlgs.Entry(ui.HotelName, "Enter the end longitude of the tour:", "")
	if err != nil {
		fmt.Println(err)
	}
	EndLongitudeFloat, err := strconv.ParseFloat(EndLongitude, 32)
	if err != nil {
		fmt.Println(err)
	}

	// GeoPoints

	EndLocation, _, err := dlgs.Entry(ui.HotelName, "Enter the end location of the tour:", "")
	if err != nil {
		fmt.Println(err)
	}

	StartTime, _, err := dlgs.Entry(ui.HotelName, "Enter the start time of the tour:", "")
	if err != nil {
		fmt.Println(err)
	}
	StartTimeTime, err := time.Parse(time.Layout, StartTime)
	if err != nil {
		fmt.Println(err)
	}

	EndTime, _, err := dlgs.Entry(ui.HotelName, "Enter the end time of the tour:", "")
	if err != nil {
		fmt.Println(err)
	}
	EndTimeTime, err := time.Parse(time.Layout, EndTime)
	if err != nil {
		fmt.Println(err)
	}

	// AcceptationState, err := dlgs.Question(ui.HotelName, "Enter the start longitude of the tour", true)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// Fill object attributes
	tour.TourName = tourName
	tour.StartLocation = StartLocation
	tour.StartLatitude = StartLatitudeFloat
	tour.EndLatitude = EndLatitudeFloat
	tour.StartLongitude = StartLongitudeFloat
	tour.EndLongitude = EndLongitudeFloat
	tour.EndLocation = EndLocation
	tour.StartTime = StartTimeTime
	tour.EndTime = EndTimeTime
	// tour.AcceptationState = AcceptationState
	tour.TourLeaderID = loggedEmployee.ID

	database.InsertTour(tour)
}

func DeleteTour() {
	tourNameToDelete := SeeAllTour()
	DeleteSelectedTour(tourNameToDelete)
	EmployeeDecider()
}

func SeeAllTour() string {
	allTours := database.GetTours()

	tourName, _, err := dlgs.List(ui.HotelName, "Please select which tour you want to delete:", allTours)
	if err != nil {
		fmt.Println(err)
	}
	return tourName
}

func DeleteSelectedTour(tourName string) {
	database.DeleteTour(tourName)
}

func AcceptTour() {
	allNewTours := database.SeeAllNewTour()
	tourName, _, err := dlgs.List(ui.HotelName, "Please select which tour you want to accept:", allNewTours)
	if err != nil {
		fmt.Println(err)
	}
	acceptSelectedTour(tourName)
}

func acceptSelectedTour(tourName string) {
	database.ApproveTour(tourName)
}
