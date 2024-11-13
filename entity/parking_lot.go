package entity

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type ParkingLot struct {
	Id            ulid.ULID `json:"id" gorm:"type:ulid;primaryKey;not null"`
	Name          string    `json:"name" form:"name" validate:"required" gorm:"not null"`
	TotalCapacity int64     `json:"totalCapacity" form:"totalCapacity" validate:"required" gorm:"not null"`
	Available     int64     `json:"available" form:"available" validate:"required" gorm:"not null"`
	Coordinate    string    `json:"coordinate" form:"coordinate" validate:"required" gorm:"unique;not null"`
	Picture       string    `json:"picture" form:"picture"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoUpdateTime;not null"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoCreateTime;not null"`
}