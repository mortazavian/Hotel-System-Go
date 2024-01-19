package logic

import (
	"fmt"

	"github.com/gen2brain/dlgs"
	ui "github.com/mortazavian/Hotel-Reservation-Go/UI"
)

func EmployeeDecider() {

	menuItems := []string{}

	if loggedEmployee.Role == "Tour Leader" {
		menuItems = append(menuItems, "Add Tour")
		menuItems = append(menuItems, "Delete Tour")
	} else if loggedEmployee.Role == "Head of Restaurant" {
		menuItems = append(menuItems, "Add Food")
		menuItems = append(menuItems, "Delete Food")
	} else if loggedEmployee.Role == "Manager" {
		menuItems = append(menuItems, "Accept Tour")
		menuItems = append(menuItems, "Add Room")
		menuItems = append(menuItems, "Add Reservation")
	} else {
		menuItems = append(menuItems, "Sign Up", "Login") //  "Manager",
		// "Head of Facilities", "Head of Reservation", "Head of Restaurant", "Head of Pool", "Tour Leader",
		// "Head of Cleaning", "Cleaner", "EXIT"))
	}

	userInput, _, err := dlgs.List(ui.HotelName, "Please select what you want.", menuItems)

	if err != nil {
		fmt.Println(err)
	}

	switch userInput {
	case "Sign Up":
		EmployeeSignUp()
	case "Login":
		EmployeeLogin()
	case "Add Tour":
		MakeNewTour()
		return
	case "Add Food":
		MakeNewFood()
		return
	case "Delete Tour":
		DeleteTour()
		return
	case "Delete Food":
		DeleteFood()
		return
	case "Accept Tour":
		AcceptTour()
		return
	case "Add Room":
		AddRoom()
		return
	case "Add Reservation":
		AddReservations()
		return
	case "Cleaner":
		// TODO
		return
	case "EXIT":
		//TODO

	}

	menuItems = nil
}
