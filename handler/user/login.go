package user

import (
	"errors"
	"github.com/Eduardo681/go_routine/helper"
	"github.com/Eduardo681/go_routine/schemas"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

// LoginHandler
// @BasePath /api/v1
// @Summary Login
// @Description Login and generate token
// @Tags User
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Request body"
// @Success 200 {object} ResponseLogin
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /login [POST]
func LoginHandler(ctx *gin.Context) {
	loginRequest := &LoginRequest{}

	if err := ctx.BindJSON(&loginRequest); err != nil {
		logger.ErrF("validation error: %v", err.Error())
		helper.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	findOne(ctx, loginRequest.Email, loginRequest.Password)
}

func findOne(ctx *gin.Context, email, password string) {
	user := &schemas.User{}

	if err := db.Where("Email = ?", email).First(user).Error; err != nil {
		helper.SendError(ctx, http.StatusBadRequest, "Invalid login credentials. Please try again")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		helper.SendError(ctx, http.StatusBadRequest, "Invalid login credentials. Please try again")
		return
	}

	tokenString, err := createToken(user)
	if err != nil {
		helper.SendError(ctx, http.StatusInternalServerError, "Failed to create token")
		return
	}

	data := &ResponseLogin{}
	data.Token = tokenString
	data.User = ConvertUserToUserResponse(user)

	helper.SendSuccess(ctx, "login", data)
}

func createToken(user *schemas.User) (string, error) {
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	tk := &schemas.Token{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)

	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
