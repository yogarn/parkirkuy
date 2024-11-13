package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/yogarn/parkirkuy/internal/handler/rest"
	"github.com/yogarn/parkirkuy/internal/repository"
	"github.com/yogarn/parkirkuy/internal/service"
	"github.com/yogarn/parkirkuy/pkg/bcrypt"
	"github.com/yogarn/parkirkuy/pkg/jwt"
	"github.com/yogarn/parkirkuy/pkg/middleware"
	"gorm.io/gorm"
)

type Config struct {
	DB  *gorm.DB
	App *fiber.App
}

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func StartUp(config *Config) {
	jwt := jwt.Init()
	bcrypt := bcrypt.Init()

	repository := repository.NewRepository(config.DB)
	service := service.NewService(repository, bcrypt, jwt)

	middleware := middleware.Init(jwt, service)

	rest := rest.NewRest(config.App, service, middleware)
	rest.RegisterRoutes()

	if os.Getenv("GO_ENV") != "test" {
		rest.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
	}
}
