package model

type VehicleDataReq struct {
	VehicleType  string `json:"vehicleType" form:"vehicleType"`
	VehicleColor string `json:"vehicleColor" form:"vehicleColor"`
	PlateNumber  string `json:"plateNumber" form:"plateNumber"`
}

type VehicleDataRes struct {
	Id           string `json:"id"`
	UserId       string `json:"userId"`
	VehicleType  string `json:"vehicleType"`
	VehicleColor string `json:"vehicleColor"`
	PlateNumber  string `json:"plateNumber"`
}
