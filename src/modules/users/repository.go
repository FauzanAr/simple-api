package users

import (
	"context"

	userentity "simple-api.com/m/src/modules/users/entities"
)

type Repository interface {
	GetUserByUsername(context.Context, string) (userentity.User, error)
	UpdateUser(context.Context, userentity.User) (error)
	RegisterUser(context.Context, userentity.User) (error)
	GetAllUsers(context.Context) ([]userentity.UserAll, error)
}