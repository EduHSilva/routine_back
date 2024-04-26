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

	api.POST("/task/rule", helper.AuthMiddleware(), task.CreateTaskRuleHandler)
	api.DELETE("/task/rule", helper.AuthMiddleware(), task.DeleteTaskRuleHandler)
	api.GET("/task/rule", helper.AuthMiddleware(), task.GetTaskRuleHandler)
	api.PUT("/task/rule", helper.AuthMiddleware(), task.UpdateTaskRuleHandler)

	api.GET("/tasks", helper.AuthMiddleware(), task.GetAllTasksRulesHandler)
	api.GET("/tasks/week", helper.AuthMiddleware(), task.GetWeekTasksHandler)
	api.PUT("/task", helper.AuthMiddleware(), task.UpdateTaskStatusHandler)

}
