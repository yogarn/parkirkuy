package service

import (
	"github.com/yogarn/parkirkuy/internal/repository"
	"github.com/yogarn/parkirkuy/pkg/bcrypt"
	"github.com/yogarn/parkirkuy/pkg/jwt"
	"github.com/yogarn/parkirkuy/pkg/ub_auth"
)

type Service struct {
	UserService       IUserService
	ParkingLotService IParkingLotService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.IBcrypt, jwt jwt.IJwt, ubAuth ub_auth.IUbAuth) *Service {
	return &Service{
		UserService:       NewUserService(repository.UserRepository, jwt, ubAuth),
		ParkingLotService: NewParkingLotService(repository.ParkingLotRepository),
	}
}
