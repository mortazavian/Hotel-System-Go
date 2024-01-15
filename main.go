package main

import (
	"fmt"

	"github.com/mortazavian/Hotel-Reservation-Go/database"
	"github.com/mortazavian/Hotel-Reservation-Go/logic"
)

func main() {
	db := database.NewGormPostgres()
	_ = db

	fmt.Println(database.Instance)

	database.MakeMigrations(db)

	logic.DecideRole()
}
