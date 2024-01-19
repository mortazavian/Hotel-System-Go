package logic

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gen2brain/dlgs"
	ui "github.com/mortazavian/Hotel-Reservation-Go/UI"
	"github.com/mortazavian/Hotel-Reservation-Go/database"

	"github.com/mortazavian/Hotel-Reservation-Go/models"
	"golang.org/x/crypto/bcrypt"
)

func TravelerSignUp() {
	MakeNewUser()
}

func MakeNewUser() {
	user := models.Traveler{}
	UserInformation(&user)
}

func UserInformation(traveler *models.Traveler) {
	firstName, _, err := dlgs.Entry(ui.HotelName, "Enter you first name:", "")
	if err != nil {
		log.Fatal(err)
	}

	lastName, _, err := dlgs.Entry(ui.HotelName, "Enter you last name:", "")
	if err != nil {
		log.Fatal(err)
	}

	username, _, err := dlgs.Entry(ui.HotelName, "Enter you username:", "")
	if err != nil {
		log.Fatal(err)
	}

	pass, _, err := dlgs.Password(ui.HotelName, "Enter you password:")
	if err != nil {
		log.Fatal(err)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(pass), 14)

	if err != nil {
		log.Fatal(err)
	}

	dataOfBirth, _, err := dlgs.Date(ui.HotelName, "Enter the date of birth:", time.Now())
	if err != nil {
		log.Fatal(err)
	}

	nationalID, _, err := dlgs.Entry(ui.HotelName, "Enter your national id:", "")
	if err != nil {
		log.Fatal(err)
	}

	phoneNumber, _, err := dlgs.Entry(ui.HotelName, "Enter your phone number: (09171234567)", "")
	if err != nil {
		log.Fatal(err)
	}

	email, _, err := dlgs.Entry(ui.HotelName, "Enter your email: (x@y.com)", "")
	if err != nil {
		log.Fatal(err)
	}

	traveler.FirstName = firstName
	traveler.LastName = lastName
	traveler.Username = username
	traveler.Password = string(password[:])
	traveler.DateofBirth = dataOfBirth
	traveler.NationalID = nationalID
	traveler.PhoneNumber = phoneNumber
	traveler.Email = email

	database.InsertTraveler(traveler)
	TravelerDecider()
}

func UserLogin() {
	username, _, err := dlgs.Entry(ui.HotelName, "enter your username: ", "")
	if err != nil {
		log.Fatal(err)
	}
	pass, _, err := dlgs.Password(ui.HotelName, "enter your password: ")
	if err != nil {
		fmt.Println(err)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		log.Fatal(err)
	}

	user, err := database.GetUser(username)
	if err != nil {
		fmt.Println("--------------------")
		dlgs.MessageBox(ui.HotelName, "Username is wrong!")
		fmt.Println(err)
		fmt.Println("--------------------")
	}
	if err == nil {
		if !CheckPasswordHash(string(password), user.Password) {
			loggedUser = user
			fmt.Println(loggedUser)
		} else {
			fmt.Println("Wrong password!!!")
		}

	}
	TravelerDecider()
}

func EmployeeLogin() {
	username, _, err := dlgs.Entry(ui.HotelName, "enter your username: ", "")
	if err != nil {
		log.Fatal(err)
	}
	pass, _, err := dlgs.Password(ui.HotelName, "enter your password: ")
	if err != nil {
		fmt.Println(err)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		log.Fatal(err)
	}

	employee, err := database.GetEmployee(username)
	if err != nil {
		fmt.Println("--------------------")
		dlgs.MessageBox(ui.HotelName, "Username is wrong!")
		fmt.Println(err)
		fmt.Println("--------------------")
	}
	if err == nil {
		if !CheckPasswordHash(string(password), employee.Password) {
			loggedEmployee = employee
			fmt.Println(loggedEmployee)
		} else {
			fmt.Println("Wrong password!!!")
		}

	}
	EmployeeDecider()
}

func EmployeeSignUp() {
	MakeNewEmployee()
}

func MakeNewEmployee() {
	employee := models.Employee{}
	EmployeeInformation(&employee)
}

func EmployeeInformation(employee *models.Employee) {
	employeeNumber, _, err := dlgs.Entry(ui.HotelName, "Enter your employee number: ", "")
	if err != nil {
		fmt.Println(err)
	}

	employeeNumberInteger, err := strconv.Atoi(employeeNumber)
	if err != nil {
		fmt.Println(err)
	}

	firstName, _, err := dlgs.Entry(ui.HotelName, "Enter your first name:", "")
	if err != nil {
		log.Fatal(err)
	}

	lastName, _, err := dlgs.Entry(ui.HotelName, "Enter your last name:", "")
	if err != nil {
		log.Fatal(err)
	}

	username, _, err := dlgs.Entry(ui.HotelName, "Enter your username:", "")
	if err != nil {
		log.Fatal(err)
	}

	pass, _, err := dlgs.Password(ui.HotelName, "Enter your password:")
	if err != nil {
		log.Fatal(err)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		log.Fatal(err)
	}

	role, _, err := dlgs.List(ui.HotelName, "Please select your role: ", []string{"Manager",
		"Head of Facilities", "Head of Reservation", "Head of Restaurant", "Head of Pool", "Tour Leader",
		"Head of Cleaning", "Cleaner", "EXIT"})
	if err != nil {
		fmt.Println(err)
	}

	employee.PersonalNumber = employeeNumberInteger
	employee.FirstName = firstName
	employee.LastName = lastName
	employee.Username = username
	employee.Password = pass
	employee.Password = string(password[:])
	employee.Role = role

	fmt.Println(employee)

	database.InsertEmployee(employee)

	EmployeeDecider()

}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AddRoom() {
	capacityList := []string{"1", "2", "4"}
	capacity, _, err := dlgs.List(ui.HotelName, "Enter the capacity of the room:", capacityList)
	if err != nil {
		fmt.Println(err)
	}

	capacityInt, err := strconv.Atoi(capacity)
	if err != nil {
		fmt.Println(err)
	}

	roomNumber, _, err := dlgs.Entry(ui.HotelName, "Enter the room number: ", "0")
	if err != nil {
		fmt.Println(err)
	}

	roomNumberInt, err := strconv.Atoi(roomNumber)
	if err != nil {
		fmt.Println(err)
	}

	roomType, _, err := dlgs.List(ui.HotelName, "Select the type of room: ", []string{"VIP", "ECO", "Penthouse"})
	if err != nil {
		fmt.Println(err)
	}

	room := models.Room{}
	room.Capacity = capacityInt
	room.RoomNumber = roomNumberInt
	room.Type = roomType

	database.InsertRoom(&room)

	EmployeeDecider()
}

func MakeNewReservation() {
	reserve := &models.Reservation{}
	FilterReservation(reserve)
}

func FilterReservation(reserve *models.Reservation) {
	roomType, _, err := dlgs.List(ui.HotelName, "Select the type of room: ", []string{"VIP", "ECO", "Penthouse"})
	if err != nil {
		fmt.Println(err)
	}

	reserveDate, _, err := dlgs.Date(ui.HotelName, "Enter the date you want room for", time.Now())
	if err != nil {
		fmt.Println(err)
	}

	roomCapacity, _, err := dlgs.List(ui.HotelName, "Please select the capacity: ", []string{"1", "2", "3"})
	if err != nil {
		fmt.Println(err)
	}

	roomCapacityInt, err := strconv.Atoi(roomCapacity)
	if err != nil {
		fmt.Println(err)
	}

	reservations := database.FilterReservationsRaw(roomCapacityInt, roomType, reserveDate)

	reservationsDates := []string{}
	for _, reserve := range reservations {
		if !reserve.GotBool {
			reservationsDates = append(reservationsDates, reserve.Date.Local().String())
		}
	}

	selectedReservation, _, err := dlgs.List(ui.HotelName, "Please select what time you want to reserve:", reservationsDates)
	if err != nil {
		fmt.Println(err)
	}

	// selectTimeReservation, err := time.Parse(selectedResevation)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	SelectReserve(selectedReservation)
}

func SelectReserve(reserveTime string) {
	database.GetReservationForUser(reserveTime, loggedUser)
}

func AddReservations() {
	date, _, err := dlgs.Date(ui.HotelName, "Enter the date: ", time.Now())
	if err != nil {
		fmt.Println(err)
	}

	roomNumber, _, err := dlgs.Entry(ui.HotelName, "Enter the room id: ", "1")
	roomNumberInt, err := strconv.Atoi(roomNumber)
	if err != nil {
		fmt.Println(err)
	}

	reservation := models.Reservation{}
	reservation.Date = date
	reservation.RoomID = uint(roomNumberInt)
	database.InsertReservation(reservation)
}
