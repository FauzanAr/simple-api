package namespacemodel

import namespaceentity "simple-api.com/m/src/modules/namespaces/entities"

type NamespaceGetDetailRequest struct {
	Id     int
	Role   string
	UserId int
}

type NamespaceGetDetailResponse struct {
	Namespace namespaceentity.Namespace `json:"namespace"`
}
