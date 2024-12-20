package service

import (
	"github.com/oklog/ulid/v2"
	"github.com/yogarn/parkirkuy/entity"
	"github.com/yogarn/parkirkuy/internal/repository"
	"github.com/yogarn/parkirkuy/model"
)

type IParkingLotService interface {
	CreateParkingLot(parkingLot *model.ParkingLotReq) (err error)
	GetParkingLotByID(id string) (parkingLot *model.ParkingLotRes, err error)
	GetParkingLotAvailableByID(id string) (available int64, err error)
	SearchParkingLotByLocation(location string) (parkingLots []*model.ParkingLotRes, err error)
	UpdateParkingLot(parkingLot *model.ParkingLotPatchReq, id string) (err error)
	DeleteParkingLot(id string) (err error)
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

func (s *ParkingLotService) GetParkingLotByID(id string) (parkingLot *model.ParkingLotRes, err error) {
	parkingLotId, err := ulid.Parse(id)
	if err != nil {
		return nil, err
	}

	parkingLotEntity, err := s.ParkingLotRepository.GetParkingLotByID(parkingLotId)
	if err != nil {
		return nil, err
	}

	parkingLot = &model.ParkingLotRes{
		Name:          parkingLotEntity.Name,
		TotalCapacity: parkingLotEntity.TotalCapacity,
		Location:      parkingLotEntity.Location,
		Coordinate:    parkingLotEntity.Coordinate,
		Picture:       parkingLotEntity.Picture,
	}

	availability, err := s.GetParkingLotAvailableByID(id)
	if err != nil {
		return nil, err
	}

	parkingLot.Available = availability

	return parkingLot, nil
}

func (s *ParkingLotService) GetParkingLotAvailableByID(id string) (available int64, err error) {
	parkingLotId, err := ulid.Parse(id)
	if err != nil {
		return -1, err
	}

	available, err = s.ParkingLotRepository.GetParkingLotAvailableByID(parkingLotId)
	if err != nil {
		return available, err
	}

	return available, nil
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
			Location:      parkingLotEntity.Location,
			Coordinate:    parkingLotEntity.Coordinate,
			Picture:       parkingLotEntity.Picture,
		}

		availability, err := s.GetParkingLotAvailableByID(parkingLotEntity.Id.String())
		if err != nil {
			return nil, err
		}

		parkingLot.Available = availability

		parkingLots = append(parkingLots, parkingLot)
	}

	return parkingLots, nil
}

func (s *ParkingLotService) UpdateParkingLot(parkingLot *model.ParkingLotPatchReq, id string) (err error) {
	parkingLotId, err := ulid.Parse(id)
	if err != nil {
		return err
	}

	parkingLotEntity := entity.ParkingLot{
		Id:            parkingLotId,
		Name:          parkingLot.Name,
		TotalCapacity: parkingLot.TotalCapacity,
		Location:      parkingLot.Location,
		Coordinate:    parkingLot.Coordinate,
		Picture:       parkingLot.Picture,
	}

	err = s.ParkingLotRepository.UpdateParkingLot(&parkingLotEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *ParkingLotService) DeleteParkingLot(id string) (err error) {
	parkingLotId, err := ulid.Parse(id)
	if err != nil {
		return err
	}

	err = s.ParkingLotRepository.DeleteParkingLot(parkingLotId)
	if err != nil {
		return err
	}

	return nil
}
