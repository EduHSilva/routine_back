package router

import (
	"github.com/Eduardo681/go_routine/docs"
	"github.com/Eduardo681/go_routine/handler/user"
	"github.com/Eduardo681/go_routine/helper"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initUserRoutes(router *gin.Engine) {
	user.InitHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	api := router.Group(basePath)

	api.POST("/login", user.LoginHandler)
	api.POST("/user", user.CreateUserHandler)
	api.GET("/user", helper.AuthMiddleware(), user.GetUserHandler)

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
