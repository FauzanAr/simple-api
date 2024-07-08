package namespaces

import (
	"context"

	namespaceentity "simple-api.com/m/src/modules/namespaces/entities"
)

type Repository interface {
	CreateNamespace(context.Context, namespaceentity.Namespace) (error)
}