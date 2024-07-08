package users

import (
	"context"

	usermodel "simple-api.com/m/src/modules/users/model"
)

type Usecase interface {
	Login(context.Context, usermodel.UserLoginRequest) (usermodel.UserLoginResponse, error)
	
}