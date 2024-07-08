package userusecase

import (
	"context"

	"simple-api.com/m/src/modules/users"
	usermodel "simple-api.com/m/src/modules/users/model"
	"simple-api.com/m/src/pkg/helper"
	"simple-api.com/m/src/pkg/logger"
	"simple-api.com/m/src/pkg/wrapper"
)

type UserUsecase struct {
	ur  users.Repository
	log logger.Logger
}

func NewUserUsecase(log logger.Logger, ur users.Repository) users.Usecase {
	return UserUsecase{
		ur:  ur,
		log: log,
	}
}

func (u UserUsecase) Login(ctx context.Context, payload usermodel.UserLoginRequest) (usermodel.UserLoginResponse, error) {
	var res usermodel.UserLoginResponse
	user, err := u.ur.GetUserByUsername(ctx, payload.Username)
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
		Status:   user.Status,
	})

	if err != nil {
		return res, wrapper.BadRequestError("Error while generating token!")
	}

	res.Token = accessToken

	return res, nil
}
