package handlers

import (
	"example/v3/db"
	"example/v3/models"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

func AppointmentURLHandler(ctx *gin.Context) {

	encodedURL := ctx.Param("encodedURL")

	db, err := db.OpenDBConnection()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Ошибка подключения к базе данных")
		return
	}

	var appointment *models.Appointments
	if err := db.Where("EncodedURL = ?", encodedURL).First(&appointment).Error; err != nil {
		ctx.String(http.StatusNotFound, "Заявка не найдена")
		return
	}

	p := bluemonday.UGCPolicy()
	p.AllowAttrs("href").OnElements("a")
	p.AllowURLSchemes("http", "https", "mailto", "javascript")

	// appointment.Date = p.Sanitize(appointment.Date)
	// appointment.Time = p.Sanitize(appointment.Time)
	// appointment.Doctor = p.Sanitize(appointment.Doctor)
	// appointment.Complain = p.Sanitize(appointment.Complain)

	ctx.HTML(http.StatusOK, "appointmentURL.html", gin.H{
		"Date":     template.HTML(appointment.Date),
		"Time":     template.HTML(appointment.Time),
		"Doctor":   template.HTML(appointment.Doctor),
		"Complain": template.HTML(appointment.Complain),
	})
}
