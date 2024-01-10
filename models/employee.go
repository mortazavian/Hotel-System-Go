package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	PersonalNumber int
	FirstName      string
	LastName       string
	Username       string
	Password       string
	Role           string
}
