package namespaceusecase

import (
	"simple-api.com/m/src/modules/namespaces"
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
