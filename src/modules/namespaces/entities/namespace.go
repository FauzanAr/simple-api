package namespaceentity

import (
	"time"
)

type Namespace struct {
	NamespaceID int       `gorm:"column:NamespaceID"`
	UserID      int       `gorm:"column:UserID"`
	TeamplateID int       `gorm:"column:TeamplateID"`
	Status      string    `gorm:"column:Status"`
	CreatedAt   time.Time `gorm:"column:CreatedAt"`
	UpdatedAt   time.Time `gorm:"column:UpdatedAt"`
}
