package usermodel

import "time"

type UserDetailRequest struct {
	Username string `omitempty`
	Id       int    `omitempty`
}

type UserDetailResponse struct {
	UserId    int64     `json:"userId"`
	Username  string    `json:"username"`
	Status    string    `json:"status"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
