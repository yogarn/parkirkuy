package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yogarn/parkirkuy/internal/service"
	"github.com/yogarn/parkirkuy/pkg/jwt"
)

type IMiddleware interface {
	AuthenticateUser(ctx *fiber.Ctx) error
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
