package router

import (
	"github.com/Eduardo681/go_routine/docs"
	"github.com/Eduardo681/go_routine/handler/opening"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initOpeningRoutes(router *gin.Engine) {
	opening.InitHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	api := router.Group(basePath)

	api.GET("/opening", opening.GetOpeningHandler)
	//api.POST("/opening", opening.CreateOpeningHandler)
	api.DELETE("/opening", opening.DeleteOpeningHandler)
	api.PUT("/opening", opening.UpdateOpeningHandler)
	api.GET("/openings", opening.GetAllOpeningHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
