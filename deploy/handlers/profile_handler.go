package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProfilePage(ctx *gin.Context) {

	isAuthenticated := ctx.GetBool("isAuthenticated")
	if isAuthenticated {
		claims, _ := ctx.Get("claims")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Profile Page",
			"claims":  claims,
		})
	} else {
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	}

}
