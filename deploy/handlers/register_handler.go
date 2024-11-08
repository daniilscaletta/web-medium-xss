package handlers

import (
	"example/v3/auth"
	"example/v3/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", gin.H{
		"Register": "Register Page",
	})
}

func RegisterHandler(ctx *gin.Context) {

	user := &models.User{
		Name:           ctx.PostForm("Name"),
		Surname:        ctx.PostForm("Surname"),
		DateOfBirthday: ctx.PostForm("DateOfBirthday"),
		Email:          ctx.PostForm("Email"),
		PhoneNumber:    ctx.PostForm("PhoneNumber"),
		Passport:       ctx.PostForm("Passport"),
		Login:          ctx.PostForm("Login"),
		Password:       ctx.PostForm("Passhash"),
	}

	err := auth.RegisterUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Redirect(http.StatusSeeOther, "/login/")
}
