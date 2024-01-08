package main

import (
	"fmt"

	"github.com/mortazavian/Hotel-Reservation-Go/database"
	"github.com/mortazavian/Hotel-Reservation-Go/decider"
	"github.com/mortazavian/Hotel-Reservation-Go/logic"
)

func main() {
	fmt.Println(logic.CheckPasswordHash("x", "$2a$14$NqXzG.V/kA3fublnf1GG.eUvGQ3kModlyXngqDEPqlzkM65bHapA2"))

	db := database.NewGormPostgres()
	_ = db

	fmt.Println(database.Instance)

	database.MakeMigrations(db)

	decider.DecideRole()
}
