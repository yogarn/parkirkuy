package entity

import "github.com/oklog/ulid/v2"

type VehicleData struct {
	Id           ulid.ULID `json:"id" gorm:"primaryKey;not null"`
	UserId       ulid.ULID `json:"userId" gorm:"not null"`
	User         User      `json:"user" gorm:"foreignKey:UserId"`
	VehicleType  string    `json:"vehicleType" form:"vehicleType" gorm:"not null"`
	VehicleColor string    `json:"vehicleColor" form:"vehicleColor" gorm:"not null"`
	PlateNumber  string    `json:"plateNumber" form:"plateNumber" gorm:"unique;not null"`
}
