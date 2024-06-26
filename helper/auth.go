package helper

import (
	"github.com/Eduardo681/go_routine/schemas"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var header = ctx.GetHeader("x-access-token")
		header = strings.TrimSpace(header)

		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing access token"})
			return
		}

		tk := &schemas.Token{}

		secret := os.Getenv("JWT_SECRET")
		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("user_id", tk.UserID)

		ctx.Set("user", tk)
		ctx.Next()
	}
}
