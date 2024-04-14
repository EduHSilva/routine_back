package opening

import (
	"github.com/Eduardo681/go_routine/helper"
	"github.com/Eduardo681/go_routine/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllOpeningHandler(ctx *gin.Context) {
	var openings []schemas.Opening

	if err := db.Find(&openings).Error; err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, "error getting openings from database")
		return
	}

	helper.SendSuccess(ctx, "list-opening", openings)
}
