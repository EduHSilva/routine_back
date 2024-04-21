package router

import (
	"github.com/Eduardo681/go_routine/docs"
	"github.com/Eduardo681/go_routine/handler/task"
	"github.com/Eduardo681/go_routine/helper"
	"github.com/gin-gonic/gin"
)

func initTaskRoutes(router *gin.Engine) {
	task.InitHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	api := router.Group(basePath)

	api.POST("/task", helper.AuthMiddleware(), task.CreateTaskHandler)
	//api.DELETE("/category", helper.AuthMiddleware(), category.DeleteCategoryHandler)
	//api.POST("/category", helper.AuthMiddleware(), category.CreateCategoryHandler)
	//api.GET("/category", helper.AuthMiddleware(), category.GetCategoryHandler)

}
