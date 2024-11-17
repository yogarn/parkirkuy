package model

import "github.com/oklog/ulid/v2"

type ParkingLotReq struct {
	Name          string `json:"name" form:"name" validate:"required"`
	TotalCapacity int64  `json:"totalCapacity" form:"totalCapacity" validate:"required"`
	Available     int64  `json:"available" form:"available" validate:"required"`
	Location      string `json:"location" form:"location" validate:"required"`
	Coordinate    string `json:"coordinate" form:"coordinate" validate:"required"`
	Picture       string `json:"picture" form:"picture"`
}

type ParkingLotPatchReq struct {
	Name          string `json:"name" form:"name"`
	TotalCapacity int64  `json:"totalCapacity" form:"totalCapacity"`
	Available     int64  `json:"available" form:"available"`
	Location      string `json:"location" form:"location"`
	Coordinate    string `json:"coordinate" form:"coordinate"`
	Picture       string `json:"picture" form:"picture"`
}

type ParkingLotRes struct {
	Id            ulid.ULID `json:"id"`
	Name          string    `json:"name"`
	TotalCapacity int64     `json:"totalCapacity"`
	Available     int64     `json:"available"`
	Location      string    `json:"location"`
	Coordinate    string    `json:"coordinate"`
	Picture       string    `json:"picture"`
}
