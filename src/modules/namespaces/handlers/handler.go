package namespacehandler

import (
	"simple-api.com/m/src/modules/namespaces"
	"simple-api.com/m/src/pkg/logger"
)

type NamespaceHandler struct {
	log logger.Logger
	nu  namespaces.Usecase
}

func NewNamespaceHandler(log logger.Logger, nu namespaces.Usecase) *NamespaceHandler {
	return &NamespaceHandler{
		log: log,
		nu:  nu,
	}
}
