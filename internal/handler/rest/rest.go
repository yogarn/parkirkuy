package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yogarn/parkirkuy/internal/service"
	"github.com/yogarn/parkirkuy/pkg/middleware"
)

type Rest struct {
	router     *fiber.App
	service    *service.Service
	middleware middleware.IMiddleware
}

func NewRest(router *fiber.App, service *service.Service, middleware middleware.IMiddleware) *Rest {
	return &Rest{
		router:     router,
		service:    service,
		middleware: middleware,
	}
}

func MountAuth(routerGroup fiber.Router, r *Rest) {
	routerGroup.Post("/login", r.LoginUBAuth)
}

func MountParkingLot(routerGroup fiber.Router, r *Rest) {
	parkingLot := routerGroup.Group("/parking-lot")
	parkingLot.Post("/", r.CreateParkingLot)
	parkingLot.Get("/:id", r.GetParkingLotByID)
	parkingLot.Get("/", r.SearchParkingLotByLocation)
	parkingLot.Patch("/:id", r.UpdateParkingLot)
	parkingLot.Delete("/:id", r.DeleteParkingLot)
}

func (r *Rest) RegisterRoutes() {
	routerGroup := r.router.Group("/api/v1")

	routerGroup.Get("/health-check", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	MountAuth(routerGroup, r)
	MountParkingLot(routerGroup, r)
}

func (r *Rest) Start(port string) error {
	return r.router.Listen(port)
}
