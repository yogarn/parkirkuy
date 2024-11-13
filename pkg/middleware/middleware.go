package middleware

import (
	"github.com/yogarn/parkirkuy/internal/service"
	"github.com/yogarn/parkirkuy/pkg/jwt"
)

type IMiddleware interface {
}

type middleware struct {
	jwtAuth jwt.IJwt
	service *service.Service
}

func Init(jwtAuth jwt.IJwt, service *service.Service) IMiddleware {
	return &middleware{
		jwtAuth: jwtAuth,
		service: service,
	}
}
