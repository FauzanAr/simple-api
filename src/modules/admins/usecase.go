package admins

import (
	"context"

	adminmodel "simple-api.com/m/src/modules/admins/models"
)

type Usecase interface {
	Login(context.Context, adminmodel.AdminLoginRequest) (adminmodel.AdminLoginResponse, error)
}