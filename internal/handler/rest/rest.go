package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yogarn/parkirkuy/internal/service"
	"github.com/yogarn/parkirkuy/pkg/middleware"
)

type Rest struct {
	router       *fiber.App
	service      *service.Service
	middleware   middleware.IMiddleware
	tradeSymbols []string
}

func NewRest(router *fiber.App, service *service.Service, middleware middleware.IMiddleware) *Rest {
	return &Rest{
		router:     router,
		service:    service,
		middleware: middleware,
	}
}

func (r *Rest) RegisterRoutes() {
	routerGroup := r.router.Group("/api/v1")

	routerGroup.Get("/health-check", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
}

func (r *Rest) Start(port string) error {
	return r.router.Listen(port)
}
