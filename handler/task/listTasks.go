package task

import (
	"github.com/Eduardo681/go_routine/helper"
	"github.com/Eduardo681/go_routine/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllTasksRulesHandler
// @BasePath /api/v1
// @Summary Get tasks
// @Description Get all tasks user
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {object} ResponseTasks
// @Failure 400 {object} helper.ErrorResponse
// @Failure 401 {object} helper.ErrorResponse
// @Security ApiKeyAuth
// @Param x-access-token header string true "Access token"
// @Router /tasks/rules [GET]
func GetAllTasksRulesHandler(ctx *gin.Context) {
	var tasksRules []schemas.TaskRule

	userID, exists := ctx.Get("user_id")
	if !exists {
		helper.SendError(ctx, http.StatusUnauthorized, "error getting tasks from database - user not found")
		return
	}

	if err := db.Where("user_id = ?", userID).Find(&tasksRules).Error; err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, "error getting tasks from database")
		return
	}

	helper.SendSuccess(ctx, "list-tasks-rules", tasksRules)
}
