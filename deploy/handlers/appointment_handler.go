package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func AppointmentPage(ctx *gin.Context) {

	isAuthenticated := ctx.GetBool("isAuthenticated")
	if isAuthenticated {
		ctx.HTML(http.StatusOK, "appointment.html", nil)
	} else {
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	}

}

func AppointmentHandler(ctx *gin.Context) {

	isAuthenticated := ctx.GetBool("isAuthenticated")
	if isAuthenticated {
		ctx.HTML(http.StatusOK, "appointment.html", nil)
	} else {
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	}

}
