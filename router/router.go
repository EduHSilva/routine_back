package router

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Init() {
	router := gin.Default()

	initRoutes(router)

	port := os.Getenv("PORT")

	if port != "" {
		err := router.Run(":" + port)
		if err != nil {
			return
		}
	} else {
		err := router.Run()
		if err != nil {
			return
		}
	}

}

func initRoutes(router *gin.Engine) {
	initUserRoutes(router)
}
