package rest

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/yogarn/parkirkuy/model"
	"github.com/yogarn/parkirkuy/pkg/response"
)

func (r *Rest) CreateParkingLot(ctx *fiber.Ctx) (err error) {
	req := new(model.ParkingLotPatchReq)
	if err = ctx.BodyParser(req); err != nil {
		return err
	}

	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		return err
	}

	err = r.service.ParkingLotService.CreateParkingLot(req)
	if err != nil {
		return err
	}

	response.Success(ctx, 200, "success", "parking lot succesfully created")
	return nil
}

func (r *Rest) GetParkingLotByID(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	parkingLot, err := r.service.ParkingLotService.GetParkingLotByID(id)
	if err != nil {
		return err
	}

	response.Success(ctx, 200, "success", parkingLot)
	return nil
}

func (r *Rest) SearchParkingLotByLocation(ctx *fiber.Ctx) (err error) {
	location := ctx.Query("location")

	parkingLots, err := r.service.ParkingLotService.SearchParkingLotByLocation(location)
	if err != nil {
		return err
	}

	response.Success(ctx, 200, "success", parkingLots)
	return nil
}

func (r *Rest) UpdateParkingLot(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")
	req := new(model.ParkingLotPatchReq)
	if err = ctx.BodyParser(req); err != nil {
		return err
	}

	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		return err
	}

	err = r.service.ParkingLotService.UpdateParkingLot(req, id)
	if err != nil {
		return err
	}

	response.Success(ctx, 200, "success", "parking lot succesfully updated")
	return nil
}

func (r *Rest) DeleteParkingLot(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	err = r.service.ParkingLotService.DeleteParkingLot(id)
	if err != nil {
		return err
	}

	response.Success(ctx, 200, "success", "parking lot succesfully deleted")
	return nil
}
