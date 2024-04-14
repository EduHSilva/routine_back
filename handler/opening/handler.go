package opening

import (
	"github.com/Eduardo681/go_routine/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("opening handler")
	db = config.GetDB()
}
