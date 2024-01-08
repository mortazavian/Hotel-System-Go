package models

import "gorm.io/gorm"

type TourLeader struct {
	gorm.Model
	PersonalNumber int
	FirstName      string
	LastName       string
	Username       string
	Password       string
	Role           string
}
