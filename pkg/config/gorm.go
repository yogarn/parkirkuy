package config

import (
	"fmt"
	"log"
	"os"

	"github.com/yogarn/parkirkuy/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func loadDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
}

func StartGorm() *gorm.DB {
	dsn := loadDSN()
	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	connection, err := db.DB()
	if err != nil {
		panic("failed to connect database")
	}

	connection.SetMaxIdleConns(10)
	connection.SetMaxOpenConns(100)

	migrate(db)
	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&entity.ParkingLot{}, &entity.User{}, &entity.Reservation{}, &entity.VehicleData{})
}
