package entity

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Reservation struct {
	Id           ulid.ULID  `json:"id" gorm:"primaryKey;not null"`
	UserId       ulid.ULID  `json:"userId" gorm:"not null"`
	User         User       `json:"user" gorm:"foreignKey:UserId"`
	ParkingLotId ulid.ULID  `json:"parkingLotId" gorm:"not null"`
	ParkingLot   ParkingLot `json:"parkingLot" gorm:"foreignKey:ParkingLotId"`
	StartTime    time.Time  `json:"startTime" gorm:"not null"`
	EndTime      time.Time  `json:"endTime" gorm:"not null"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"autoUpdateTime;not null"`
	UpdatedAt    time.Time  `json:"updatedAt" gorm:"autoCreateTime;not null"`
}
