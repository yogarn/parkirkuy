package service

import (
	"github.com/oklog/ulid/v2"
	"github.com/yogarn/parkirkuy/entity"
	"github.com/yogarn/parkirkuy/internal/repository"
	"github.com/yogarn/parkirkuy/model"
)

type IParkingLotService interface {
	CreateParkingLot(parkingLot *model.ParkingLotReq) (err error)
	GetParkingLotByID(id string) (parkingLot *model.ParkingLotReq, err error)
	SearchParkingLotByLocation(location string) (parkingLots []*model.ParkingLotRes, err error)
}

type ParkingLotService struct {
	ParkingLotRepository repository.IParkingLotRepository
}

func NewParkingLotService(parkingLotRepository repository.IParkingLotRepository) IParkingLotService {
	return &ParkingLotService{
		ParkingLotRepository: parkingLotRepository,
	}
}

func (s *ParkingLotService) CreateParkingLot(parkingLot *model.ParkingLotReq) (err error) {
	parkingLotEntity := entity.ParkingLot{
		Id:            ulid.Make(),
		Name:          parkingLot.Name,
		TotalCapacity: parkingLot.TotalCapacity,
		Available:     parkingLot.Available,
		Location:      parkingLot.Location,
		Coordinate:    parkingLot.Coordinate,
		Picture:       parkingLot.Picture,
	}

	err = s.ParkingLotRepository.CreateParkingLot(&parkingLotEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *ParkingLotService) GetParkingLotByID(id string) (parkingLot *model.ParkingLotReq, err error) {
	parkingLotId, err := ulid.Parse(id)
	if err != nil {
		return nil, err
	}

	parkingLotEntity, err := s.ParkingLotRepository.GetParkingLotByID(parkingLotId)
	if err != nil {
		return nil, err
	}

	parkingLot = &model.ParkingLotReq{
		Name:          parkingLotEntity.Name,
		TotalCapacity: parkingLotEntity.TotalCapacity,
		Available:     parkingLotEntity.Available,
		Location:      parkingLotEntity.Location,
		Coordinate:    parkingLotEntity.Coordinate,
		Picture:       parkingLotEntity.Picture,
	}

	return parkingLot, nil
}

func (s *ParkingLotService) SearchParkingLotByLocation(location string) (parkingLots []*model.ParkingLotRes, err error) {
	parkingLotEntities, err := s.ParkingLotRepository.SearchParkingLotByLocation(location)
	if err != nil {
		return nil, err
	}

	for _, parkingLotEntity := range parkingLotEntities {
		parkingLot := &model.ParkingLotRes{
			Id:            parkingLotEntity.Id,
			Name:          parkingLotEntity.Name,
			TotalCapacity: parkingLotEntity.TotalCapacity,
			Available:     parkingLotEntity.Available,
			Location:      parkingLotEntity.Location,
			Coordinate:    parkingLotEntity.Coordinate,
			Picture:       parkingLotEntity.Picture,
		}

		parkingLots = append(parkingLots, parkingLot)
	}

	return parkingLots, nil
}