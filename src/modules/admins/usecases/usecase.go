package adminusecase

import (
	"context"

	"simple-api.com/m/src/modules/admins"
	adminmodel "simple-api.com/m/src/modules/admins/models"
	"simple-api.com/m/src/pkg/logger"
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

	return res, nil
}
