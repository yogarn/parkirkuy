package repository

import (
	"github.com/oklog/ulid/v2"
	"github.com/yogarn/parkirkuy/entity"
	"github.com/yogarn/parkirkuy/pkg/response"
	"gorm.io/gorm"
)

type IParkingLotRepository interface {
	CreateParkingLot(parkingLot *entity.ParkingLot) (err error)
	GetParkingLotByID(id ulid.ULID) (parkingLot *entity.ParkingLot, err error)
	SearchParkingLotByLocation(location string) (parkingLots []*entity.ParkingLot, err error)
}

type ParkingLotRepository struct {
	db *gorm.DB
}

func NewParkingLotRepository(db *gorm.DB) IParkingLotRepository {
	return &ParkingLotRepository{
		db: db,
	}
}

func (r *ParkingLotRepository) CreateParkingLot(parkingLot *entity.ParkingLot) (err error) {
	err = r.db.Create(parkingLot).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ParkingLotRepository) GetParkingLotByID(id ulid.ULID) (parkingLot *entity.ParkingLot, err error) {
	parkingLot = new(entity.ParkingLot)
	err = r.db.Where("id = ?", id).First(parkingLot).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &response.ParkingLotNotFound
		}
		return nil, err
	}

	return parkingLot, nil
}

func (r *ParkingLotRepository) SearchParkingLotByLocation(location string) (parkingLots []*entity.ParkingLot, err error) {
	err = r.db.Where("location LIKE ?", "%"+location+"%").Find(&parkingLots).Error
	if err != nil {
		return nil, err
	}

	return parkingLots, nil
}
