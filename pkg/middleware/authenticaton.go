package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/yogarn/parkirkuy/pkg/response"
)

func (m *middleware) AuthenticateUser(ctx *fiber.Ctx) error {
	bearer := ctx.Get("Authorization")

	if bearer == "" {
		return &response.InvalidToken
	}

	token := strings.Split(bearer, " ")[1]
	userId, err := m.jwtAuth.ValidateToken(token)
	if err != nil {
		return &response.InvalidToken
	}

	user, err := m.service.UserService.GetUserById(userId.String())
	if err != nil {
		return err
	}

	ctx.Locals("user", user)
	return ctx.Next()
}
