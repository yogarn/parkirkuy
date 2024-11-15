package config

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/yogarn/parkirkuy/pkg/response"

	u "github.com/ahmdyaasiin/ub-auth-without-notification/v2"
)

func StartFiber() *fiber.App {
	app := fiber.New(
		fiber.Config{
			ErrorHandler: CustomErrorHandler,
		},
	)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept",
		AllowCredentials: false,
	}))

	return app
}

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	var errorRequest *response.ErrorRequest
	if errors.As(err, &errorRequest) {
		code = errorRequest.Code
		message = errorRequest.Error()
	}

	var ubAuthError *u.ResponseDetails
	if errors.As(err, &ubAuthError) {
		code = ubAuthError.Code
		message = ubAuthError.Message
	}

	var fiberError *fiber.Error
	if errors.As(err, &fiberError) {
		code = fiberError.Code
	}

	var validationError validator.ValidationErrors
	if errors.As(err, &validationError) {
		code = fiber.StatusBadRequest
		message = "Bad Request"
	}

	response.Error(ctx, code, message, err)

	return nil
}
