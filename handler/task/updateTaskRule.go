package task

import (
	"fmt"
	"github.com/Eduardo681/go_routine/helper"
	"github.com/Eduardo681/go_routine/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateTaskRuleHandler
// @BasePath /api/v1
// @Summary Update task rule
// @Description Update task rule info
// @Tags Task
// @Accept json
// @Produce json
// @Param id query uint64 true "id"
// @Param request body UpdateTaskRequest true "Request body"
// @Success 200 {object} ResponseTasks
// @Failure 400 {object} helper.ErrorResponse
// @Failure 401 {object} helper.ErrorResponse
// @Security ApiKeyAuth
// @Param x-access-token header string true "Access token"
// @Router /user [PUT]
func UpdateTaskRuleHandler(ctx *gin.Context) {
	request := UpdateTaskRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		logger.ErrF("validation error: %v", err.Error())
		helper.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.ErrF("validation error: %v", err.Error())
		helper.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")

	if id == "" {
		helper.SendError(ctx, http.StatusBadRequest,
			helper.ErrParamIsRequired("id", "query param").Error())
		return
	}

	taskRule := &schemas.TaskRule{}

	if err := db.First(&taskRule, id).Error; err != nil {
		helper.SendError(ctx, http.StatusNotFound, fmt.Sprintf("taskRule with id %s not found", id))
		return
	}

	if request.Title != "" {
		taskRule.Title = request.Title
	}

	if request.Priority != "" {
		taskRule.Priority = schemas.Priority(request.Priority)
	}

	if request.CategoryID != 0 {
		taskRule.CategoryID = request.CategoryID
	}

	if err := db.Save(&taskRule).Error; err != nil {
		logger.ErrF("error updating: %s", err.Error())
		helper.SendError(ctx, http.StatusInternalServerError, err.Error())
	}

	helper.SendSuccess(ctx, "update-task-rule", taskRule)
}
