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

func (n NamespaceUsecase) CreateNamespace(ctx context.Context, payload namespacemodel.NamespaceCreateRequest) error {
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

func (n NamespaceUsecase) DeleteNamespace(ctx context.Context, payload namespacemodel.NamespaceDeleteRequest) error {
	err := n.nr.DeleteNamespace(ctx, payload.Id)
	if err != nil {
		return err
	}

	return nil
}

func (n NamespaceUsecase) GetAllNamespaces(ctx context.Context) (namespacemodel.NamespaceGetAllResponse, error) {
	var res namespacemodel.NamespaceGetAllResponse

	namespaces, err := n.nr.GetAllNamespaces(ctx)
	if err != nil {
		return res, err
	}

	res.Namespaces = namespaces

	return res, nil
}
