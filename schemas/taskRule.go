package schemas

import (
	"gorm.io/gorm"
	"time"
)

type TaskRule struct {
	gorm.Model
	Title      string
	Frequency  Frequency
	Priority   Priority
	DateStart  time.Time
	DateEnd    time.Time
	StartTime  time.Time
	EndTime    time.Time
	CategoryID uint     `gorm:"not null"`
	Category   Category `gorm:"foreignKey:CategoryID"`
	UserID     uint     `gorm:"not null"`
	User       Category `gorm:"foreignKey:UserID"`
}
