package adminentity

import "time"

type Admin struct {
	UserID       int       `gorm:"coloumn:UserId"`
	Username     string    `gorm:"coloumn:Username"`
	PasswordHash string    `gorm:"coloumn:PasswordHash"`
	Role         string    `gorm:"coloumn:Role"`
	Email        string    `gorm:"coloumn:Email"`
	CreatedAt    time.Time `gorm:"column:CreatedAt"`
	UpdatedAt    time.Time `gorm:"column:UpdatedAt"`
}
