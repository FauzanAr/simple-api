package usermodel

type UserRegisterRequest struct {
	Username     string `json:"username" validate:"required"`
	PasswordHash string `json:"passwordHash" validate:"required"`
	Email        string `json:"email" validate:"required"`
}
