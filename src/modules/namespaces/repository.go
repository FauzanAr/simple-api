package namespaces

import (
	"context"

	namespaceentity "simple-api.com/m/src/modules/namespaces/entities"
)

type Repository interface {
	CreateNamespace(context.Context, namespaceentity.Namespace) (error)
	DeleteNamespace(context.Context, int) (error)
	GetAllNamespaces(context.Context) ([]namespaceentity.Namespace, error)
	GetNamespaceById(context.Context, int) (namespaceentity.Namespace, error)
}