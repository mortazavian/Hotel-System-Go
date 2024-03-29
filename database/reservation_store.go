package database

import (
	"fmt"
	"strings"
	"time"

	"github.com/mortazavian/Hotel-Reservation-Go/models"
)

func FilterReservationsRaw(capacity int, roomType string, date time.Time) []models.Reservation {
	var reservations []models.Reservation
	Instance.Raw("select * from reservations join rooms on reservations.room_id = rooms.id where rooms.type = ? and reservations.date != ? and rooms.capacity = ?", roomType, date, capacity).Scan(&reservations)
	return reservations
}

func InsertReservation(reservation models.Reservation) {
	fmt.Println(reservation.Date)
	Instance.Save(&reservation)
}

func GetReservationForUser(reserveTime string, traveler models.Traveler) {
	reserveTime = strings.TrimRight(reserveTime, "+330")

	// myDate, err := time.Parse(time.RFC3339, reserveTime)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// reservation := models.Reservation{Date: myDate}
	// _ = reservation

	Instance.Model(&models.Reservation{}).Where("date = ?", reserveTime).Update("user_id", traveler.ID).Update("got_bool", true)
}

func DeleteReservation(reserveTime string) {
	reserveTime = strings.TrimRight(reserveTime, "+330")

	// reserve := models.Reservation{}

	Instance.Delete(&models.Reservation{}, "date = ?", reserveTime)
}

func UpdateReservationEndDate(reserveID int, newDate time.Time) {

	newDateString := newDate.String()

	fmt.Println(newDateString)

	newDateString = strings.TrimRight(newDateString, "UTC")

	fmt.Println(newDateString)

	newDateString = strings.TrimRight(newDateString, "0000+")

	fmt.Println(newDateString)

	// fmt.Println(newDate.Format("2024-01-19 19:53:56.204"))

	Instance.Model(&models.Reservation{}).Where("id = ?", reserveID).Update("end_date", newDateString)
}
