package userentity

import "time"

type UserAll struct {
	UserID    uint      `gorm:"column:UserId"`
	Username  string    `gorm:"column:Username"`
	Status    string    `gorm:"column:Status"`
	Email     string    `gorm:"column:Email"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt"`
}

func (UserAll) TableName() string {
	return "users"
}
