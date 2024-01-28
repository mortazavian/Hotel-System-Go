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

func AddPoolSans() {
	sans := CreateNewPoolSans()
	PoolSansInformation(sans)
}

func CreateNewPoolSans() *models.PoolSans {
	return &models.PoolSans{}
}

func PoolSansInformation(sans *models.PoolSans) {
	startTime, _, err := dlgs.Date(ui.HotelName, "Enter the start time of the sans: ", time.Now())
	if err != nil {
		fmt.Println(err)
	}

	endTime, _, err := dlgs.Date(ui.HotelName, "Enter the start end of the sans: ", time.Now())
	if err != nil {
		fmt.Println(err)
	}

	capacity, _, err := dlgs.Entry(ui.HotelName, "Please enter the capacity of the sans", "50")
	if err != nil {
		fmt.Println(err)
	}

	capacityInteger, err := strconv.Atoi(capacity)
	if err != nil {
		fmt.Println(err)
	}

	sans.Capacity = capacityInteger
	sans.StartTime = startTime
	sans.EndTime = endTime

	database.InsertPoolSans(sans)
	EmployeeDecider()

}

func SeePoolSans() {
	sanses := database.GetAllSans()
	sansesString := []string{}
	for _, sans := range sanses {
		// startTimeString := string.
		sansesString = append(sansesString, sans.StartTime.String()[0:11]+" - "+sans.EndTime.String()[0:11])

	}

	_, _, err := dlgs.List(ui.HotelName, "All sans. 8:00 A.M. to 11:00 P.M.", sansesString)
	if err != nil {
		fmt.Println(err)
	}

	TravelerDecider()

}
