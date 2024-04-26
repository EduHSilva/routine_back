package task

import (
	"fmt"
	"github.com/Eduardo681/go_routine/helper"
	"github.com/Eduardo681/go_routine/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// DeleteTaskRuleHandler
// @BasePath /api/v1
// @Summary Delete a task
// @Description Delete a task and yours status
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {object} ResponseTask
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Failure 401 {object} helper.ErrorResponse
// @Param id query uint64 true "id"
// @Security ApiKeyAuth
// @Param x-access-token header string true "Access token"
// @Router /task/rule [DELETE]
func DeleteTaskRuleHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		helper.SendError(ctx, http.StatusBadRequest,
			helper.ErrParamIsRequired("id", "query param").Error())
		return
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			logger.Err("recovered in CreateTaskHandler")
		}
	}()

	taskRule := &schemas.TaskRule{}

	if err := db.First(&taskRule, id).Error; err != nil {
		helper.SendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id %s not found", id))
		return
	}

	if err := deleteTaskStatus(tx, id); err != nil {
		tx.Rollback()
		helper.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	tx.Commit()

	helper.SendSuccess(ctx, "delete-taskRule", taskRule)
}

func deleteTaskStatus(tx *gorm.DB, id string) error {
	tx.Delete(&schemas.Task{}, "taskID = ?", id)
	return nil
}
