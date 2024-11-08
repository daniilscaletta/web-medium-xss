package middleware

import (
	"example/v3/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authTokenString := ctx.GetHeader("Authorization")
		if authTokenString == "" {
			token, err := ctx.Cookie("Authorization")
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"error": "Missing Authorization Header",
				})
				ctx.Abort()
				return
			}
			authTokenString = token
		}

		tokenString := authTokenString
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
