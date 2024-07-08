package adminentity

import "time"

type Admin struct {
	UserID       int       `gorm:"column:UserId"`
	Username     string    `gorm:"column:Username"`
	PasswordHash string    `gorm:"column:PasswordHash"`
	Role         string    `gorm:"column:Role"`
	Email        string    `gorm:"column:Email"`
	CreatedAt    time.Time `gorm:"column:CreatedAt"`
	UpdatedAt    time.Time `gorm:"column:UpdatedAt"`
}
