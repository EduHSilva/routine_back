package schemas

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	LastLogin time.Time
}
