package usermodel

import "time"

type UserUpdateRequest struct {
	Username         string `json:"username" validate:"required"`
	Status           string `json:"status" validate:"required"`
	Email            string `json:"email" validate:"required"`
	OriginalUsername string `omitempty`
	Id               int    `omitempty`
}

type UserUpdateResponse struct {
	UserId    int64     `json:"userId"`
	Username  string    `json:"username"`
	Status    string    `json:"status"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
