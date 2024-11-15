package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository       IUserRepository
	ParkingLotRepository IParkingLotRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:       NewUserRepository(db),
		ParkingLotRepository: NewParkingLotRepository(db),
	}
}
