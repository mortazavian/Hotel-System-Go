package database

import (
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var once sync.Once

// var db *gorm.DB

func NewGormPostgres() *gorm.DB {
	once.Do(func() {
		tehranTimezone, _ := time.LoadLocation("Asia/Tehran")

		// Connection configuration
		dsn := &url.URL{
			Scheme:   "postgres",
			User:     url.UserPassword("root", "PGIVo79HWrGRcbwxVi9D4uXk"),
			Host:     "kilimanjaro.liara.cloud:30996",
			Path:     "hotel_reservation_go",
			RawQuery: "timezone=" + tehranTimezone.String(),
		}

		// Convert URL to connection string
		connStr := dsn.String()

		db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to the database:", err)
		}

		fmt.Println("Successfully connected to the database!")

		Instance = db
	})

	return Instance
}
