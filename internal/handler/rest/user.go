package rest

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/yogarn/parkirkuy/model"
	"github.com/yogarn/parkirkuy/pkg/response"
)

func (r *Rest) LoginUBAuth(ctx *fiber.Ctx) (err error) {
	req := new(model.UBAuthReq)
	if err = ctx.BodyParser(req); err != nil {
		return err
	}

	validate := validator.New()
	if err = validate.Struct(req); err != nil {
		return err
	}

	resp, err := r.service.UserService.LoginUB(req.Identifier, req.Password)
	if err != nil {
		return err
	}

	response.Success(ctx, 200, "success", &model.JwtRes{Token: resp})
	return nil
}
