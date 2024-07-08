package namespacemodel

import namespaceentity "simple-api.com/m/src/modules/namespaces/entities"

type NamespaceGetAllResponse struct {
	Namespaces []namespaceentity.Namespace `json:"namespaces"`
}
