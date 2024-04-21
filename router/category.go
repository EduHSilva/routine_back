package router

import (
	"github.com/Eduardo681/go_routine/docs"
	"github.com/Eduardo681/go_routine/handler/category"
	"github.com/Eduardo681/go_routine/helper"
	"github.com/gin-gonic/gin"
)

func initCategoryRoutes(router *gin.Engine) {
	category.InitHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	api := router.Group(basePath)

	api.GET("/category", helper.AuthMiddleware(), category.GetCategoryHandler)
	api.GET("/categories", helper.AuthMiddleware(), category.GetAllCategoriesHandler)
	api.DELETE("/category", helper.AuthMiddleware(), category.DeleteCategoryHandler)
	api.POST("/category", helper.AuthMiddleware(), category.CreateCategoryHandler)

}
