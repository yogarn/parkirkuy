package rest

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/yogarn/parkirkuy/model"
	"github.com/yogarn/parkirkuy/pkg/response"
)

func (r *Rest) CreateVehicleData(ctx *fiber.Ctx) (err error) {
	req := new(model.VehicleDataReq)
	if err = ctx.BodyParser(req); err != nil {
		return err
	}

	userIdString := (ctx.Locals("user").(*model.UserRes).Id)

	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		return err
	}

	err = r.service.VehicleDataService.CreateVehicleData(userIdString, req)
	if err != nil {
		return err
	}

	response.Success(ctx, 201, "success", "vehicle data created")
	return nil
}

func (r *Rest) GetVehicleDataByUserId(ctx *fiber.Ctx) (err error) {
	userId := (ctx.Locals("user").(*model.UserRes).Id)

	vehicleData, err := r.service.VehicleDataService.GetVehicleDataByUserId(userId)
	if err != nil {
		return err
	}

	response.Success(ctx, 200, "success", vehicleData)
	return nil
}

func (r *Rest) GetVehicleDataById(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")
	vehicleData, err := r.service.VehicleDataService.GetVehicleDataById(id)
	if err != nil {
		return err
	}

	response.Success(ctx, 200, "success", vehicleData)
	return nil
}
