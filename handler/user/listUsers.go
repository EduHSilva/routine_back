package user

import (
	"github.com/Eduardo681/go_routine/helper"
	"github.com/Eduardo681/go_routine/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsersHandler(ctx *gin.Context) {
	var users []schemas.User

	if err := db.Find(&users).Error; err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, "error getting users from database")
		return
	}

	helper.SendSuccess(ctx, "list-users", users)
}
