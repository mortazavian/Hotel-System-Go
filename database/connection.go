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

var instance *gorm.DB
var once sync.Once

func NewGormPostgres() *gorm.DB {
	once.Do(func() {
		tehranTimezone, _ := time.LoadLocation("Asia/Tehran")

		// Connection configuration
		dsn := &url.URL{
			Scheme:   "postgres",
			User:     url.UserPassword("postgres", "admin"),
			Host:     "localhost",
			Path:     "hotel_reservation_go",
			RawQuery: "sslmode=disable&timezone=" + tehranTimezone.String(),
		}

		// Convert URL to connection string
		connStr := dsn.String()

		db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to the database:", err)
		}

		fmt.Println("Successfully connected to the database!")

		instance = db
	})

	return instance
}
