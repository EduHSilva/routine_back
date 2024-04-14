package user

import (
	"github.com/Eduardo681/go_routine/helper"
	"github.com/Eduardo681/go_routine/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUserHandler
// @BasePath /api/v1
// @Summary Create user
// @Description Create a new User
// @Tags User
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "Request body"
// @Success 200 {object} ResponseUser
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /user [POST]
func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}
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

	hash, err := helper.HashPassword(request.Password)
	if err != nil {
		logger.ErrF("hash error: %v", err.Error())
	}

	user := schemas.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hash,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.ErrF("Error in CreateUserHandler: %s", err.Error())
		helper.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	helper.SendSuccess(ctx, "create-user", user)
}
