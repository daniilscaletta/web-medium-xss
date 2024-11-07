package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	tmpl, _ := template.ParseFiles("templates/home.html")
// 	tmpl.Execute(w, nil)

// }

func HomeHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", nil)
}
