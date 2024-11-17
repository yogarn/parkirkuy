package repository

import (
	"github.com/oklog/ulid/v2"
	"github.com/yogarn/parkirkuy/entity"
	"github.com/yogarn/parkirkuy/pkg/response"
	"gorm.io/gorm"
)

type IReservationRepository interface {
	CreateReservation(reservation *entity.Reservation) (err error)
	GetReservationByID(id ulid.ULID) (reservation *entity.Reservation, err error)
	GetReservationByUserID(userID ulid.ULID) (reservations []*entity.Reservation, err error)
	GetReservationByParkingLotID(parkingLotID ulid.ULID) (reservations []*entity.Reservation, err error)
	UpdateReservation(reservation *entity.Reservation) (err error)
	DeleteReservation(id ulid.ULID) (err error)
}

type ReservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) IReservationRepository {
	return &ReservationRepository{
		db: db,
	}
}

func (r *ReservationRepository) CreateReservation(reservation *entity.Reservation) (err error) {
	err = r.db.Create(reservation).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ReservationRepository) GetReservationByID(id ulid.ULID) (reservation *entity.Reservation, err error) {
	reservation = new(entity.Reservation)
	err = r.db.Where("id = ?", id).First(reservation).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &response.ReservationNotFound
		}
		return nil, err
	}

	return reservation, nil
}

func (r *ReservationRepository) GetReservationByUserID(userID ulid.ULID) (reservations []*entity.Reservation, err error) {
	err = r.db.Where("user_id = ?", userID).Find(&reservations).Error
	if err != nil {
		return nil, err
	}

	return reservations, nil
}

func (r *ReservationRepository) GetReservationByParkingLotID(parkingLotID ulid.ULID) (reservations []*entity.Reservation, err error) {
	err = r.db.Where("parking_lot_id = ?", parkingLotID).Find(&reservations).Error
	if err != nil {
		return nil, err
	}

	return reservations, nil
}

func (r *ReservationRepository) UpdateReservation(reservation *entity.Reservation) (err error) {
	res := r.db.Updates(reservation)

	if res.RowsAffected == 0 {
		return &response.ReservationNotFound
	}

	err = res.Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ReservationRepository) DeleteReservation(id ulid.ULID) (err error) {
	res := r.db.Where("id = ?", id).Delete(&entity.Reservation{})
	if res.RowsAffected == 0 {
		return &response.ReservationNotFound
	}

	err = res.Error
	if err != nil {
		return err
	}

	return nil
}
