package middlewares

import (
	"net/http"

	"example.com/api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	authToken := context.Request.Header.Get("Authorization")
	if authToken == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	userId, err := utils.VerifyToken(authToken)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	context.Set("userID", userId)

	context.Next()
}
