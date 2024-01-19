package models

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Capacity   int
	RoomNumber int
	Type       string
}
