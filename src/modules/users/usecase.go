package users

import (
	"context"

	usermodel "simple-api.com/m/src/modules/users/model"
)

type Usecase interface {
	Login(context.Context, usermodel.UserLoginRequest) (usermodel.UserLoginResponse, error)
	GetUserDetail(context.Context, usermodel.UserDetailRequest) (usermodel.UserDetailResponse, error)
	UpdateUser(context.Context, usermodel.UserUpdateRequest) (usermodel.UserUpdateResponse, error)
	RegisterUser(context.Context, usermodel.UserRegisterRequest) (error)
}