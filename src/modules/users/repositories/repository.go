package userrepository

import (
	"context"
	"strings"

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

func (u UserRepository) UpdateUser(ctx context.Context, user userentity.User) (error) {
	err := u.db.GetDatabase().Where("UserId = ?", user.UserID).Save(&user).Error
	if err != nil {
		u.log.Error(ctx, err.Error(), err, nil)

		if strings.Contains(err.Error(), "Duplicate entry") {
			return wrapper.BadRequestError("Duplicate entry for email or username")
		}

		return wrapper.InternalServerError("Error while saving data!")
	}

	return nil
}

func (u UserRepository) RegisterUser(ctx context.Context, user userentity.User) (error) {
	err := u.db.GetDatabase().Create(&user).Error
	if err != nil {
		u.log.Error(ctx, err.Error(), err, nil)

		if strings.Contains(err.Error(), "Duplicate entry") {
			return wrapper.BadRequestError("Duplicate entry for email or username")
		}

		return wrapper.InternalServerError("Error while saving data!")
	}

	return nil
}

func (u UserRepository) GetAllUsers(ctx context.Context) ([]userentity.UserAll, error) {
	var result []userentity.UserAll
	err := u.db.GetDatabase().Find(&result).Error
	if err != nil {
		u.log.Error(ctx, err.Error(), err, nil)
		if err.Error() == "record not found" {
			return result, wrapper.NotFoundError("User not found!")
		}

		return result, wrapper.InternalServerError("Error while read data!")
	}

	return result, nil
}

func (u UserRepository) GetUserById(ctx context.Context, id int) (userentity.User, error) {
	var result userentity.User
	err := u.db.GetDatabase().Where("UserId = ?", id).First(&result).Error
	if err != nil {
		u.log.Error(ctx, err.Error(), err, nil)
		if err.Error() == "record not found" {
			return result, wrapper.NotFoundError("User not found!")
		}

		return result, wrapper.InternalServerError("Error while read data!")
	}

	return result, nil
}
