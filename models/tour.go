package models

import (
	"time"

	"gorm.io/gorm"
)

type Tour struct {
	gorm.Model
	TourName       string
	StartLocation  string
	StartLatitude  float64
	StartLongitude float64
	EndLatitude    float64
	EndLongitude   float64
	GeoPoints      float32
	EndLocation    string
	StartTime      time.Time
	EndTime        time.Time
	AcceptionState bool
}
