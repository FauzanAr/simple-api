package namespacemodel

type NamespaceGetStatusRequest struct {
	Id int
}

type NamespaceGetStatusResponse struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}
