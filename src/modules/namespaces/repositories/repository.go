package namespacerepository

import (
	"context"

	"simple-api.com/m/src/modules/namespaces"
	namespaceentity "simple-api.com/m/src/modules/namespaces/entities"
	"simple-api.com/m/src/pkg/databases/mysql"
	"simple-api.com/m/src/pkg/logger"
	"simple-api.com/m/src/pkg/wrapper"
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

func (nr NamespaceRepository) CreateNamespace(ctx context.Context, namespace namespaceentity.Namespace) error {
	err := nr.db.GetDatabase().Create(&namespace).Error
	if err != nil {
		nr.log.Error(ctx, err.Error(), err, nil)

		return wrapper.InternalServerError("Error while saving data!")
	}

	return nil
}

func (nr NamespaceRepository) DeleteNamespace(ctx context.Context, id int) (error) {
	err := nr.db.GetDatabase().Delete(&namespaceentity.Namespace{}, id).Error
	if err != nil {
		nr.log.Error(ctx, err.Error(), err, nil)

		return wrapper.InternalServerError("Error while delete data!")
	}

	return nil
}
