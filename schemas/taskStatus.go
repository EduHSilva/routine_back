package schemas

import (
	"gorm.io/gorm"
	"time"
)

type TaskStatus struct {
	gorm.Model
	Status StatusTask
	Date   time.Time
	TaskID uint `gorm:"not null"`
}
