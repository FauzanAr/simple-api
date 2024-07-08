package userrepository

import (
	"context"

	"simple-api.com/m/src/modules/users"
	userentity "simple-api.com/m/src/modules/users/entities"
	"simple-api.com/m/src/pkg/databases/mysql"
	"simple-api.com/m/src/pkg/logger"
	"simple-api.com/m/src/pkg/wrapper"
)

type UserRepository struct {
	db  *mysql.Mysql
	log logger.Logger
}

func NewUserRepository(log logger.Logger, db *mysql.Mysql) users.Repository {
	return UserRepository{
		db:  db,
		log: log,
	}
}

func (u UserRepository) GetUserByUsername(ctx context.Context, username string) (userentity.User, error) {
	var result userentity.User
	err := u.db.GetDatabase().Where("username = ?", username).First(&result).Error
	if err != nil {
		u.log.Error(ctx, err.Error(), err, nil)
		if err.Error() == "record not found" {
			return result, wrapper.NotFoundError("User not found!")
		}

		return result, wrapper.InternalServerError("Error while read data!")
	}

	return result, nil
}
