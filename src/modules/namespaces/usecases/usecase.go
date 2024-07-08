package namespaceusecase

import (
	"context"

	"simple-api.com/m/src/modules/namespaces"
	namespaceentity "simple-api.com/m/src/modules/namespaces/entities"
	namespacemodel "simple-api.com/m/src/modules/namespaces/models"
	"simple-api.com/m/src/pkg/logger"
)

type NamespaceUsecase struct {
	nr  namespaces.Repository
	log logger.Logger
}

func NewNamespaceUsecase(log logger.Logger, nr namespaces.Repository) namespaces.Usecase {
	return NamespaceUsecase{
		nr:  nr,
		log: log,
	}
}

func (n NamespaceUsecase) CreateNamespace(ctx context.Context, payload namespacemodel.NamespaceCreateRequest) (error) {
	var namespace namespaceentity.Namespace

	namespace.UserID = payload.UserID
	namespace.TeamplateID = payload.TeamplateID
	namespace.Status = payload.Status

	err := n.nr.CreateNamespace(ctx, namespace)
	if err != nil {
		return err
	}

	return nil
}
