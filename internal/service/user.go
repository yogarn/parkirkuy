package service

import (
	"errors"

	"github.com/oklog/ulid/v2"
	"github.com/yogarn/parkirkuy/entity"
	"github.com/yogarn/parkirkuy/internal/repository"
	"github.com/yogarn/parkirkuy/pkg/jwt"
	"github.com/yogarn/parkirkuy/pkg/response"
	"github.com/yogarn/parkirkuy/pkg/ub_auth"
)

type IUserService interface {
	LoginUB(identifier string, password string) (jwtToken string, err error)
}

type UserService struct {
	Jwt    jwt.IJwt
	UBAuth ub_auth.IUbAuth

	UserRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository, jwt jwt.IJwt, ubAuth ub_auth.IUbAuth) IUserService {
	return &UserService{
		UserRepository: userRepository,

		Jwt:    jwt,
		UBAuth: ubAuth,
	}
}

func (s *UserService) LoginUB(identifier string, password string) (jwtToken string, err error) {
	resp, err := s.UBAuth.Login(identifier, password)
	if err != nil {
		return "", err
	}

	user, err := s.UserRepository.GetUserByEmail(resp.Email)
	if err != nil || user == nil {
		if !errors.Is(err, &response.UserNotFound) {
			return "", err
		}

		addUser := &entity.User{
			Id:       ulid.Make(),
			Name:     resp.FullName,
			Username: resp.NIM,
			Email:    resp.Email,
		}

		err = s.UserRepository.CreateUser(addUser)
		if err != nil {
			return "", err
		}
	}

	newUser, err := s.UserRepository.GetUserByEmail(resp.Email)
	if err != nil {
		return "", err
	}

	token, err := s.Jwt.CreateToken(newUser.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}
