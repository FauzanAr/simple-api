package namespacemodel

type NamespaceCreateRequest struct {
	TeamplateID int    `json:"teamplateId" validate:"required"`
	Status      string `json:"status" validate:"required"`
	UserID      int    `json:"userId" validate:"required"`
}
