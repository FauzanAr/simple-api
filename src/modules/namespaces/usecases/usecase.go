package namespaceusecase

import (
	"context"

	"simple-api.com/m/src/modules/namespaces"
	namespaceentity "simple-api.com/m/src/modules/namespaces/entities"
	namespacemodel "simple-api.com/m/src/modules/namespaces/models"
	"simple-api.com/m/src/pkg/logger"
	"simple-api.com/m/src/pkg/wrapper"
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

func (n NamespaceUsecase) GetStatusNamespace(ctx context.Context, payload namespacemodel.NamespaceGetStatusRequest) (namespacemodel.NamespaceGetStatusResponse, error) {
	var res namespacemodel.NamespaceGetStatusResponse
	namespace, err := n.nr.GetNamespaceById(ctx, payload.Id)
	if err != nil {
		return res, err
	}

	res.Id = namespace.NamespaceID
	res.Status = namespace.Status

	return res, nil
}

func (n NamespaceUsecase) GetDetailNamespace(ctx context.Context, payload namespacemodel.NamespaceGetDetailRequest) (namespacemodel.NamespaceGetDetailResponse, error) {
	var res namespacemodel.NamespaceGetDetailResponse
	if payload.Role != "ADMIN" && payload.Role != "USER" {
		return res, wrapper.UnauthorizedError("Only user and admin can view this namespace")
	}

	namespace, err := n.nr.GetNamespaceById(ctx, payload.Id)
	if err != nil {
		return res, err
	}

	if payload.Role == "USER" && namespace.UserID != payload.UserId {
		return res, wrapper.UnauthorizedError("Only admin can view this namespace")
	}

	res.Namespace = namespace

	return res, nil
}
