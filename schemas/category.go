package schemas

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title string `gorm:"embedded;not null"`
}
