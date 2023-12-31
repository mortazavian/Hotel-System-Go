package models

import (
	"time"

	"gorm.io/gorm"
)

type Traveler struct {
	gorm.Model
	FirstName   string
	LastName    string
	Username    string
	Password    string
	DateofBirth time.Time
	NationalID  string
	phoneNumber string
	email       string
}
