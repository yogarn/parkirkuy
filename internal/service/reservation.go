package service

import (
	"github.com/oklog/ulid/v2"
	"github.com/yogarn/parkirkuy/entity"
	"github.com/yogarn/parkirkuy/internal/repository"
	"github.com/yogarn/parkirkuy/model"
)

type IReservationService interface {
	CreateReservation(userIdString string, reservation *model.ReservationReq) (err error)
	GetReservationByID(idString string) (reservation *model.ReservationRes, err error)
	GetReservationByUserID(userIdString string) (reservations []*model.ReservationRes, err error)
	GetReservationByParkingLotID(parkingLotIdString string) (reservations []*model.ReservationRes, err error)
	UpdateReservation(idString string, reservation *model.ReservationPatchReq) (err error)
	DeleteReservation(idString string) (err error)
}

type ReservationService struct {
	ReservationRepository repository.IReservationRepository
}

func NewReservationService(reservationRepository repository.IReservationRepository) IReservationService {
	return &ReservationService{
		ReservationRepository: reservationRepository,
	}
}

func (s *ReservationService) CreateReservation(userIdString string, reservation *model.ReservationReq) (err error) {
	userId, err := ulid.Parse(userIdString)
	if err != nil {
		return err
	}

	reservationEntity := entity.Reservation{
		Id:           ulid.Make(),
		UserId:       userId,
		ParkingLotId: reservation.ParkingLotId,
		StartTime:    reservation.StartTime,
		EndTime:      reservation.EndTime,
	}

	err = s.ReservationRepository.CreateReservation(&reservationEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *ReservationService) GetReservationByID(idString string) (reservation *model.ReservationRes, err error) {
	id, err := ulid.Parse(idString)
	if err != nil {
		return nil, err
	}

	reservationEntity, err := s.ReservationRepository.GetReservationByID(id)
	if err != nil {
		return nil, err
	}

	reservation = &model.ReservationRes{
		Id:           reservationEntity.Id,
		UserId:       reservationEntity.UserId,
		ParkingLotId: reservationEntity.ParkingLotId,
		StartTime:    reservationEntity.StartTime,
		EndTime:      reservationEntity.EndTime,
		CreatedAt:    reservationEntity.CreatedAt,
		UpdatedAt:    reservationEntity.UpdatedAt,
	}

	return reservation, nil
}

func (s *ReservationService) GetReservationByUserID(userIdString string) (reservations []*model.ReservationRes, err error) {
	userId, err := ulid.Parse(userIdString)
	if err != nil {
		return nil, err
	}

	reservationEntities, err := s.ReservationRepository.GetReservationByUserID(userId)
	if err != nil {
		return nil, err
	}

	for _, reservationEntity := range reservationEntities {
		reservation := &model.ReservationRes{
			Id:           reservationEntity.Id,
			UserId:       reservationEntity.UserId,
			ParkingLotId: reservationEntity.ParkingLotId,
			StartTime:    reservationEntity.StartTime,
			EndTime:      reservationEntity.EndTime,
			CreatedAt:    reservationEntity.CreatedAt,
			UpdatedAt:    reservationEntity.UpdatedAt,
		}

		reservations = append(reservations, reservation)
	}

	return reservations, nil
}

func (s *ReservationService) GetReservationByParkingLotID(parkingLotIdString string) (reservations []*model.ReservationRes, err error) {
	parkingLotId, err := ulid.Parse(parkingLotIdString)
	if err != nil {
		return nil, err
	}

	reservationEntities, err := s.ReservationRepository.GetReservationByParkingLotID(parkingLotId)
	if err != nil {
		return nil, err
	}

	for _, reservationEntity := range reservationEntities {
		reservation := &model.ReservationRes{
			Id:           reservationEntity.Id,
			UserId:       reservationEntity.UserId,
			ParkingLotId: reservationEntity.ParkingLotId,
			StartTime:    reservationEntity.StartTime,
			EndTime:      reservationEntity.EndTime,
			CreatedAt:    reservationEntity.CreatedAt,
			UpdatedAt:    reservationEntity.UpdatedAt,
		}

		reservations = append(reservations, reservation)
	}

	return reservations, nil
}

func (s *ReservationService) UpdateReservation(idString string, reservation *model.ReservationPatchReq) (err error) {
	id, err := ulid.Parse(idString)
	if err != nil {
		return err
	}

	reservationEntity := entity.Reservation{
		Id:           id,
		ParkingLotId: reservation.ParkingLotId,
		StartTime:    reservation.StartTime,
		EndTime:      reservation.EndTime,
	}

	err = s.ReservationRepository.UpdateReservation(&reservationEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *ReservationService) DeleteReservation(id string) (err error) {
	err = s.ReservationRepository.DeleteReservation(ulid.MustParse(id))
	if err != nil {
		return err
	}

	return nil
}
