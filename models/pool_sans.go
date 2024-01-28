package models

import (
	"time"

	"gorm.io/gorm"
)

type PoolSans struct {
	gorm.Model
	StartTime time.Time
	EndTime   time.Time
	Capacity  int
}
