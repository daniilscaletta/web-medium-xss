package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProfileHandler(ctx *gin.Context) {

	claims, _ := ctx.Get("claims")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Profile Page",
		"claims":  claims,
	})

}
