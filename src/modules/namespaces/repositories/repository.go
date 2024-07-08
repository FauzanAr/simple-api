package namespacerepository

import (
	"simple-api.com/m/src/modules/namespaces"
	"simple-api.com/m/src/pkg/databases/mysql"
	"simple-api.com/m/src/pkg/logger"
)

type NamespaceRepository struct {
	db  *mysql.Mysql
	log logger.Logger
}

func NewNamespaceRepository(log logger.Logger, db *mysql.Mysql) namespaces.Repository {
	return NamespaceRepository{
		db:  db,
		log: log,
	}
}
