package admins

import (
	"context"

	adminentity "simple-api.com/m/src/modules/admins/entities"
)

type Repository interface {
	GetAdminByUsername(context.Context, string) (adminentity.Admin, error)
}