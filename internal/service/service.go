package service

import (
	"github.com/yogarn/parkirkuy/internal/repository"
	"github.com/yogarn/parkirkuy/pkg/bcrypt"
	"github.com/yogarn/parkirkuy/pkg/jwt"
)

type Service struct {
}

func NewService(repository *repository.Repository, bcrypt bcrypt.IBcrypt, jwt jwt.IJwt) *Service {
	return &Service{}
}
