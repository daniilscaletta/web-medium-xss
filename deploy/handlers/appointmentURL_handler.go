package handlers

import (
	"example/v3/db"
	"example/v3/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AppointmentURLHandler(ctx *gin.Context) {

	isAuthenticated := ctx.GetBool("isAuthenticated")
	if isAuthenticated {

		encodedURL := ctx.Param("encodedURL")

		db, err := db.OpenDBConnection()
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Ошибка подключения к базе данных")
			return
		}

		var appointment models.Appointments
		if err := db.Where("EncodedURL = ?", encodedURL).First(&appointment).Error; err != nil {
			ctx.String(http.StatusNotFound, "Заявка не найдена")
			return
		}

		ctx.HTML(http.StatusOK, "appointmentURL.html", gin.H{
			"Date":     appointment.Date,
			"Time":     appointment.Time,
			"Doctor":   appointment.Doctor,
			"Complain": appointment.Complain,
		})

	} else {
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
	}
}
