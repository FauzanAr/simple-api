package namespaces

import (
	"context"

	namespacemodel "simple-api.com/m/src/modules/namespaces/models"
)

type Usecase interface {
	CreateNamespace(context.Context, namespacemodel.NamespaceCreateRequest) (error)
	DeleteNamespace(context.Context, namespacemodel.NamespaceDeleteRequest) (error)
	GetAllNamespaces(context.Context) (namespacemodel.NamespaceGetAllResponse,error)
}