package usermodel

import "time"

type UserDetailRequest struct {
	UserID int64
}

type UserDetailResponse struct {
	UserId    string    `json:"userId"`
	Username  string    `json:"username"`
	Status    string    `json:"status"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
