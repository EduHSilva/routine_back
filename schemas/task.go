package schemas

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Title      string
	Frequency  Frequency
	Priority   Priority
	Date       time.Time
	StartTime  time.Time
	EndTime    time.Time
	CategoryID uint `gorm:"not null"`
	UserID     uint `gorm:"not null"`
}
