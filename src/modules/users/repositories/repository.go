package userrepository

import (
	"simple-api.com/m/src/modules/users"
	"simple-api.com/m/src/pkg/databases/mysql"
	"simple-api.com/m/src/pkg/logger"
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
