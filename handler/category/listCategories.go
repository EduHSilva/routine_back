package category

import (
	"github.com/Eduardo681/go_routine/helper"
	"github.com/Eduardo681/go_routine/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllCategoriesHandler
// @BasePath /api/v1
// @Summary Get categories
// @Description Get all categories
// @Tags Category
// @Accept json
// @Produce json
// @Success 200 {object} ResponseCategories
// @Failure 400 {object} helper.ErrorResponse
// @Router /categories [GET]
func GetAllCategoriesHandler(ctx *gin.Context) {
	var categories []schemas.Category

	if err := db.Find(&categories).Error; err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, "error getting categories from database")
		return
	}

	helper.SendSuccess(ctx, "list-categories", categories)
}
