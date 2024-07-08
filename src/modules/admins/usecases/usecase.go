package adminusecase

import (
	"context"

	"simple-api.com/m/src/modules/admins"
	adminmodel "simple-api.com/m/src/modules/admins/models"
	"simple-api.com/m/src/pkg/helper"
	"simple-api.com/m/src/pkg/logger"
	"simple-api.com/m/src/pkg/wrapper"
)

type AdminUsecase struct {
	ar  admins.Repository
	log logger.Logger
}

func NewAdminUsecase(log logger.Logger, ar admins.Repository) admins.Usecase {
	return AdminUsecase{
		ar:  ar,
		log: log,
	}
}

func (a AdminUsecase) Login(ctx context.Context, payload adminmodel.AdminLoginRequest) (adminmodel.AdminLoginResponse, error) {
	var res adminmodel.AdminLoginResponse

	user, err := a.ar.GetAdminByUsername(ctx, payload.Username)
	if err != nil {
		return res, err
	}
	isPasswordSame, err := helper.Compare(payload.Password, user.PasswordHash)
	if err != nil {
		return res, wrapper.InternalServerError(err.Error())
	}

	if !isPasswordSame {
		return res, wrapper.BadRequestError("Password incorrect")
	}

	accessToken, err := helper.GenerateAccessToken(ctx, helper.Claims{
		Id:       int64(user.UserID),
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}, "ADMIN")

	if err != nil {
		return res, wrapper.BadRequestError("Error while generating token!")
	}

	res.Token = accessToken
	return res, nil
}
