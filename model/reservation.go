package model

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type ReservationReq struct {
	ParkingLotId ulid.ULID `json:"parkingLotId" validate:"required"`
	StartTime    time.Time `json:"startTime" validate:"required"`
	EndTime      time.Time `json:"endTime" validate:"required"`
}

type ReservationPatchReq struct {
	ParkingLotId ulid.ULID `json:"parkingLotId"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
}

type ReservationRes struct {
	Id           ulid.ULID `json:"id"`
	UserId       ulid.ULID `json:"userId"`
	ParkingLotId ulid.ULID `json:"parkingLotId"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
