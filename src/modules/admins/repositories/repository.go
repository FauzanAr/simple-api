package adminrepository

import (
	"context"

	"simple-api.com/m/src/modules/admins"
	adminentity "simple-api.com/m/src/modules/admins/entities"
	"simple-api.com/m/src/pkg/databases/mysql"
	"simple-api.com/m/src/pkg/logger"
	"simple-api.com/m/src/pkg/wrapper"
)

type AdminRepository struct {
	db  *mysql.Mysql
	log logger.Logger
}

func NewAdminRepository(log logger.Logger, db *mysql.Mysql) admins.Repository {
	return AdminRepository{
		db:  db,
		log: log,
	}
}

func (a AdminRepository) GetAdminByUsername(ctx context.Context, username string) (adminentity.Admin, error) {
	var result adminentity.Admin
	err := a.db.GetDatabase().Where("username = ?", username).First(&result).Error
	if err != nil {
		a.log.Error(ctx, err.Error(), err, nil)
		if err.Error() == "record not found" {
			return result, wrapper.NotFoundError("Admin not found!")
		}

		return result, wrapper.InternalServerError("Error while read data!")
	}

	return result, nil
}
