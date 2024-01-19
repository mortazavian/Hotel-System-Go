package models

import "gorm.io/gorm"

type CleaningRequest struct {
	gorm.Model
	UserID        uint
	Traveler      Traveler `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ReservationID uint
	Reservation   Reservation `gorm:"foreignKey:ReservationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
