package models

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Name             string
	Ingredient       string
	Detail           string
	RestaurantHeadID uint
}
