package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func RecordPage(ctx *gin.Context) {

	isAuthenticated := ctx.GetBool("isAuthenticated")
	if isAuthenticated {
		ctx.JSON(http.StatusOK, gin.H{
			"Record":   "Recording Page",
			"Date":     "17.11",
			"Time":     "10:00",
			"Location": "New York",
			"Doctor":   "Tooth Dother",
		})
	} else {
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	}

}

func RecordHandler(ctx *gin.Context) {

	isAuthenticated := ctx.GetBool("isAuthenticated")
	if isAuthenticated {
		ctx.JSON(http.StatusOK, gin.H{
			"Record":   "Recording Page",
			"Date":     "17.11",
			"Time":     "10:00",
			"Location": "New York",
			"Doctor":   "Tooth Dother",
		})
	} else {
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	}

}
