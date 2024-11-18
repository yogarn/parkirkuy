package repository

import (
	"github.com/oklog/ulid/v2"
	"github.com/yogarn/parkirkuy/entity"
	"gorm.io/gorm"
)

type IVehicleDataRepository interface {
	CreateVehicleData(*entity.VehicleData) (err error)
	GetVehicleDataById(id ulid.ULID) (*entity.VehicleData, error)
	GetVehicleDataByUserId(userId ulid.ULID) (*[]entity.VehicleData, error)
}

type VehicleDataRepository struct {
	db *gorm.DB
}

func NewVehicleDataRepository(db *gorm.DB) IVehicleDataRepository {
	return &VehicleDataRepository{db}
}

func (r *VehicleDataRepository) CreateVehicleData(vehicleData *entity.VehicleData) (err error) {
	err = r.db.Create(vehicleData).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *VehicleDataRepository) GetVehicleDataById(id ulid.ULID) (*entity.VehicleData, error) {
	var vehicleData entity.VehicleData
	err := r.db.Where("id = ?", id).First(&vehicleData).Error
	if err != nil {
		return nil, err
	}
	return &vehicleData, nil
}

func (r *VehicleDataRepository) GetVehicleDataByUserId(userId ulid.ULID) (*[]entity.VehicleData, error) {
	var vehicleData []entity.VehicleData
	err := r.db.Where("user_id = ?", userId).Find(&vehicleData).Error
	if err != nil {
		return nil, err
	}
	return &vehicleData, nil
}
