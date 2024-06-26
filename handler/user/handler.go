package user

import (
	"github.com/Eduardo681/go_routine/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("handler user")
	db = config.GetDB()
}
