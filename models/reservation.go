package models

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	Date     time.Time
	RoomID   uint
	Room     Room `gorm:"foreignKey:RoomID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	GotBool  bool
	UserID   uint
	Traveler Traveler `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
