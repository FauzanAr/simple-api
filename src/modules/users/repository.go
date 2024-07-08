package users

import (
	"context"

	userentity "simple-api.com/m/src/modules/users/entities"
)

type Repository interface {
	GetUserByUsername(context.Context, string) (userentity.User, error)
}