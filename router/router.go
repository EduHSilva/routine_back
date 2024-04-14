package router

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	initRoutes(router)

	err := router.Run()
	if err != nil {
		return

	}
}

func initRoutes(router *gin.Engine) {
	initUserRoutes(router)
}
