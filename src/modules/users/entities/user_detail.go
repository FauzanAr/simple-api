package userentity

import "time"

type User struct {
	UserID       uint      `gorm:"column:UserId"`
	Username     string    `gorm:"column:Username"`
	PasswordHash string    `gorm:"column:PasswordHash"`
	Status       string    `gorm:"column:Status"`
	Email        string    `gorm:"column:Email"`
	CreatedAt    time.Time `gorm:"column:CreatedAt"`
	UpdatedAt    time.Time `gorm:"column:UpdatedAt"`
}
