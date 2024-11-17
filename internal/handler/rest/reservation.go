package rest

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/yogarn/parkirkuy/model"
	"github.com/yogarn/parkirkuy/pkg/response"
)

func (r *Rest) CreateReservation(ctx *fiber.Ctx) (err error) {
	req := new(model.ReservationReq)
	userId := (ctx.Locals("user").(*model.UserRes).Id)
	if err = ctx.BodyParser(req); err != nil {
		return err
	}

	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		return err
	}

	err = r.service.ReservationService.CreateReservation(userId, req)
	if err != nil {
		return err
	}

	response.Success(ctx, 201, "success", "reservation created")
	return nil
}

func (r *Rest) GetReservations(ctx *fiber.Ctx) (err error) {
	userId := (ctx.Locals("user").(*model.UserRes).Id)
	reservations, err := r.service.ReservationService.GetReservationByUserID(userId)
	if err != nil {
		return err
	}

	response.Success(ctx, 200, "success", reservations)
	return nil
}

func (r *Rest) GetReservation(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")
	reservation, err := r.service.ReservationService.GetReservationByID(id)
	if err != nil {
		return err
	}

	response.Success(ctx, 200, "success", reservation)
	return nil
}

func (r *Rest) GetReservationByParkingLot(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")
	reservations, err := r.service.ReservationService.GetReservationByParkingLotID(id)
	if err != nil {
		return err
	}

	response.Success(ctx, 200, "success", reservations)
	return nil
}

func (r *Rest) UpdateReservation(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")
	req := new(model.ReservationPatchReq)
	if err = ctx.BodyParser(req); err != nil {
		return err
	}

	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		return err
	}

	err = r.service.ReservationService.UpdateReservation(id, req)
	if err != nil {
		return err
	}

	response.Success(ctx, 200, "success", "reservation updated")
	return nil
}

func (r *Rest) DeleteReservation(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")
	err = r.service.ReservationService.DeleteReservation(id)
	if err != nil {
		return err
	}

	response.Success(ctx, 200, "success", "reservation deleted")
	return nil
}
