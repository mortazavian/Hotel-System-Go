package main

import (
	"fmt"

	"github.com/gen2brain/dlgs"
	"github.com/mortazavian/Hotel-Reservation-Go/database"
)

func main() {
	fmt.Println("HEllo")
	db := database.NewGormPostgres()
	_ = db

	database.MakeMigrations(db)
	item, _, err := dlgs.List("List", "Select item from list:", []string{"Bug", "New Feature", "Improvement"})
	if err != nil {
		panic(err)
	}
	fmt.Println(item)
}
