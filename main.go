package main

import (
	"fmt"

	"github.com/mortazavian/Hotel-Reservation-Go/database"
	"github.com/mortazavian/Hotel-Reservation-Go/logic"
	"gorm.io/gorm"
)

func main() {

	var db *gorm.DB
	// go func() {
	db = database.NewGormPostgres()
	// }()
	// for {
	// 	dlgs.MessageBox(ui.HotelName, "Trying to connect to Database!!!")

	// 	fmt.Println("sdfsfs")
	// 	if db != nil {
	// 		break
	// 	}
	// }

	fmt.Println(database.Instance)

	database.MakeMigrations(db)

	logic.DecideRole()
}

// Added features:
// 1. Head of pool can add a sans.
// 2. Traveler can see pool sanses.
