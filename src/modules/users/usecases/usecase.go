package userusecase

import (
	"simple-api.com/m/src/modules/users"
	"simple-api.com/m/src/pkg/logger"
)

type UserUsecase struct {
	ur  users.Usecase
	log logger.Logger
}

func NewUserUsecase(log logger.Logger, ur users.Repository) users.Usecase {
	return UserUsecase{
		ur:  ur,
		log: log,
	}
}
