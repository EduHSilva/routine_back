package schemas

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title  string `gorm:"embedded;not null" json:"title"`
	UserID uint   `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserID" json:"user"`
}
