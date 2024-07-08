package userusecase

import (
	"context"

	"simple-api.com/m/src/modules/users"
	userentity "simple-api.com/m/src/modules/users/entities"
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
	}, "USER")

	if err != nil {
		return res, wrapper.BadRequestError("Error while generating token!")
	}

	res.Token = accessToken

	return res, nil
}

func (u UserUsecase) GetUserDetail(ctx context.Context, payload usermodel.UserDetailRequest) (usermodel.UserDetailResponse, error) {
	var res usermodel.UserDetailResponse
	user, err := u.ur.GetUserByUsername(ctx, payload.Username)
	if err != nil {
		return res, err
	}

	res.UserId = int64(user.UserID)
	res.Username = user.Username
	res.Status = user.Status
	res.Email = user.Email
	res.CreatedAt = user.CreatedAt
	res.UpdatedAt = user.UpdatedAt

	return res, nil
}

func (u UserUsecase) UpdateUser(ctx context.Context, payload usermodel.UserUpdateRequest) (usermodel.UserUpdateResponse, error) {
	var res usermodel.UserUpdateResponse
	user, err := u.ur.GetUserByUsername(ctx, payload.OriginalUsername)
	if err != nil {
		return res, err
	}

	user.Username = payload.Username
	user.Status = payload.Status
	user.Email = payload.Email

	err = u.ur.UpdateUser(ctx, user)
	if err != nil {
		return res, err
	}

	res.UserId = int64(user.UserID)
	res.Username = user.Username
	res.Email = user.Email
	res.Status = user.Status
	res.CreatedAt = user.CreatedAt
	res.UpdatedAt = user.UpdatedAt

	return res, nil
}

func (u UserUsecase) RegisterUser(ctx context.Context, payload usermodel.UserRegisterRequest) error {
	var user userentity.User

	hashedPassword, err := helper.Hash(payload.PasswordHash)
	if err != nil {
		return wrapper.BadRequestError("Error while hashing password")
	}

	user.Username = payload.Username
	user.Email = payload.Email
	user.PasswordHash = hashedPassword
	user.Status = "Status" // Change this to enum

	err = u.ur.RegisterUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
