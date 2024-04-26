package task

import (
	"github.com/Eduardo681/go_routine/helper"
	"github.com/Eduardo681/go_routine/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// GetWeekTasksHandler
// @BasePath /api/v1
// @Summary Get tasks of week
// @Description A map that associates dates (in string format) with ResponseDateTasks.
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {object} ResponseTaskMap
// @Failure 400 {object} helper.ErrorResponse
// @Failure 401 {object} helper.ErrorResponse
// @Security ApiKeyAuth
// @Param x-access-token header string true "Access token"
// @Param currentDate query string true "currentDate"
// @Router /tasks/week [GET]
func GetWeekTasksHandler(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")

	if !exists {
		helper.SendError(ctx, http.StatusUnauthorized, "error getting tasks from database - user not found")
		return
	}

	currentDateStr := ctx.Query("currentDate")

	currentDate, err := time.Parse(time.DateOnly, currentDateStr)
	if err != nil {
		currentDate = time.Now()
	}

	query := getTaskStatusQuery(userID.(uint), currentDate)

	var task []schemas.Task
	if err := query.Find(&task).Error; err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, "error getting tasks from database")
		return
	}

	taskMap := make(ResponseTaskMap)
	for _, ts := range task {
		truncatedDate := ts.Date.Truncate(24 * time.Hour).Format(time.DateOnly)

		if _, ok := taskMap[truncatedDate]; !ok {
			taskMap[truncatedDate] = &ResponseDateTasks{
				Tasks: make([]ResponseDataWeekTask, 0),
			}
		}

		taskMap[truncatedDate].Tasks = append(taskMap[truncatedDate].Tasks, ResponseDataWeekTask{
			ID:           ts.ID,
			IDTask:       ts.TaskRuleID,
			Title:        ts.TaskRule.Title,
			Priority:     string(ts.TaskRule.Priority),
			StartTime:    formatTime(ts.TaskRule.StartTime),
			EndTime:      formatTime(ts.TaskRule.EndTime),
			CategoryName: ts.TaskRule.Category.Title,
			Done:         ts.Done,
		})
	}

	helper.SendSuccess(ctx, "get-week-tasks", taskMap)
}

func formatTime(date time.Time) string {
	return date.Format(time.TimeOnly)
}

func getTaskStatusQuery(userID uint, currentDate time.Time) *gorm.DB {
	query := db.Model(&schemas.Task{})

	query = query.Select("tasks.id, task_id, t.category_id, t.title, status, tasks.date, t.start_time, t.end_time, t.priority, c.title")

	query = query.Joins("INNER JOIN tasks t ON t.id = task_id")
	query = query.Joins("INNER JOIN categories c ON t.category_id = c.id")

	query = query.Where("t.user_id = ?", userID)
	query = query.Where("tasks.date between ? and ?", currentDate, currentDate.AddDate(0, 0, 7))

	query = query.Preload("TaskRule").Preload("TaskRule.Category")
	return query
}
