package entity

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Reservation struct {
	Id           ulid.ULID  `json:"id" gorm:"type:ulid;primaryKey;not null"`
	UserId       ulid.ULID  `json:"userId" gorm:"type:ulid;not null"`
	User         User       `json:"user" gorm:"foreignKey:UserId"`
	ParkingLotId ulid.ULID  `json:"parkingLotId" gorm:"type:ulid;not null"`
	ParkingLot   ParkingLot `json:"parkingLot" gorm:"foreignKey:ParkingLotId"`
	StartTime    time.Time  `json:"startTime" gorm:"not null"`
	EndTime      time.Time  `json:"endTime" gorm:"not null"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"autoUpdateTime;not null"`
	UpdatedAt    time.Time  `json:"updatedAt" gorm:"autoCreateTime;not null"`
}
