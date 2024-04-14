package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func Init() {
	router := gin.Default()

	initRoutes(router)

	port := os.Getenv("PORT")
	addr := fmt.Sprintf("localhost:%s", port)
	err := router.Run(addr)
	if err != nil {
		return
	}
}

func initRoutes(router *gin.Engine) {
	initUserRoutes(router)
}
