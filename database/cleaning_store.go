package database

import (
	"fmt"
	"strings"

	"github.com/mortazavian/Hotel-Reservation-Go/models"
)

func GetAllUserReservations(traveler models.Traveler) []models.Reservation {
	reservations := []models.Reservation{}
	Instance.Where("user_id = ?", traveler.ID).Find(&reservations)
	return reservations
}

func InsertCleaningRequest(reserveTime string, traveler models.Traveler) {
	reserveTime = strings.TrimRight(reserveTime, "+330")

	var reservation models.Reservation

	Instance.Where("date = ?", reserveTime).First(&reservation)

	var cleaningRequest models.CleaningRequest
	cleaningRequest.UserID = traveler.ID
	cleaningRequest.ReservationID = reservation.ID

	result := Instance.Create(&cleaningRequest)
	_ = result
}

func GetAllCleaningRequests() []models.CleaningRequest {
	cleaningRequests := []models.CleaningRequest{}

	result := Instance.Find(&cleaningRequests)
	_ = result

	for i, _ := range cleaningRequests {
		if cleaningRequests[i].Done {
			remove(cleaningRequests, i)
		}
	}

	return cleaningRequests
}

func remove(slice []models.CleaningRequest, s int) []models.CleaningRequest {
	return append(slice[:s], slice[s+1:]...)
}

func GetAllWorkers() []models.Employee {

	workers := []models.Employee{}
	result := Instance.Find(&models.Employee{}).Where("role = ?", "Cleaner").Find(&workers)
	_ = result
	return workers
}

func AssignCleaningToWorker(workerId, reservationId string) {
	fmt.Println("+++++++++++++++++")
	fmt.Println(workerId, reservationId)
	fmt.Println("+++++++++++++++++")

	Instance.Model(&models.CleaningRequest{}).Where("reservation_id = ?", reservationId).Update("worker_id", workerId)
}
