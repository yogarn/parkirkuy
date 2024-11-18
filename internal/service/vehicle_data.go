package service

import (
	"github.com/oklog/ulid/v2"
	"github.com/yogarn/parkirkuy/entity"
	"github.com/yogarn/parkirkuy/internal/repository"
	"github.com/yogarn/parkirkuy/model"
)

type IVehicleDataService interface {
	CreateVehicleData(userIdString string, vehicleReq *model.VehicleDataReq) (err error)
	GetVehicleDataById(idString string) (vehicleData *model.VehicleDataRes, err error)
	GetVehicleDataByUserId(userIdString string) (vehicleData []*model.VehicleDataRes, err error)
}

type VehicleDataService struct {
	VehicleDataRepository repository.IVehicleDataRepository
}

func NewVehicleDataService(vehicleDataRepository repository.IVehicleDataRepository) IVehicleDataService {
	return &VehicleDataService{
		VehicleDataRepository: vehicleDataRepository,
	}
}

func (s *VehicleDataService) CreateVehicleData(userIdString string, vehicleData *model.VehicleDataReq) (err error) {
	userId, err := ulid.Parse(userIdString)
	if err != nil {
		return err
	}

	vehicle := &entity.VehicleData{
		Id:           ulid.Make(),
		UserId:       userId,
		VehicleType:  vehicleData.VehicleType,
		VehicleColor: vehicleData.VehicleColor,
		PlateNumber:  vehicleData.PlateNumber,
	}

	err = s.VehicleDataRepository.CreateVehicleData(vehicle)
	if err != nil {
		return err
	}

	return nil
}

func (s *VehicleDataService) GetVehicleDataById(idString string) (vehicleData *model.VehicleDataRes, err error) {
	id, err := ulid.Parse(idString)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	vehicleDataEntity, err := s.VehicleDataRepository.GetVehicleDataById(id)
	if err != nil {
		return nil, err
	}

	vehicleData = &model.VehicleDataRes{
		Id:           vehicleDataEntity.Id.String(),
		UserId:       vehicleDataEntity.UserId.String(),
		VehicleType:  vehicleDataEntity.VehicleType,
		VehicleColor: vehicleDataEntity.VehicleColor,
		PlateNumber:  vehicleDataEntity.PlateNumber,
	}

	return vehicleData, nil
}

func (s *VehicleDataService) GetVehicleDataByUserId(userIdString string) (vehicleDatas []*model.VehicleDataRes, err error) {
	userId, err := ulid.Parse(userIdString)
	if err != nil {
		return nil, err
	}

	vehicleEntities, err := s.VehicleDataRepository.GetVehicleDataByUserId(userId)
	if err != nil {
		return nil, err
	}

	for _, vehicleEntity := range *vehicleEntities {
		vehicle := &model.VehicleDataRes{
			Id:           vehicleEntity.Id.String(),
			UserId:       vehicleEntity.UserId.String(),
			VehicleType:  vehicleEntity.VehicleType,
			VehicleColor: vehicleEntity.VehicleColor,
			PlateNumber:  vehicleEntity.PlateNumber,
		}

		vehicleDatas = append(vehicleDatas, vehicle)
	}

	return vehicleDatas, nil
}
