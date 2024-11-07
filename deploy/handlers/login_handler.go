package handlers

import (
	"example/v3/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func LoginPage(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "login.html", gin.H{
		"Login": "Login Page",
	})
}
func LoginHandler(ctx *gin.Context) {

	login := ctx.PostForm("Login")
	password := ctx.PostForm("Password")

	user, err := auth.AuthenticateUser(login, password)
	if err != nil {
		ctx.JSON(http.StatusUnavailableForLegalReasons, gin.H{
			"error": "Invalid login or password",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":        "Logged in successfully",
		"name":           user.Name,
		"surname":        user.Surname,
		"email":          user.Email,
		"passport":       user.Passport,
		"dateOfBirthday": user.DateOfBirthday,
		"phoneNumber":    user.PhoneNumber,
		"login":          user.Login,
		"passHash":       user.PassHash,
	})

}

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		login := r.FormValue("Login")
// 		password := r.FormValue("Password")

// 		user, err := auth.AuthenticateUser(login, password)
// 		if err != nil {
// 			http.Error(w, "error auth: "+err.Error(), http.StatusUnauthorized)
// 			return
// 		}

// 		fmt.Fprintf(w, "welcome, %s!\n", user.Name)
// 		fmt.Fprintf(w, "user Surname, %s\n", user.Surname)
// 		fmt.Fprintf(w, "user Passport, %s\n", user.Passport)
// 		fmt.Fprintf(w, "user Email, %s\n", user.Email)
// 		fmt.Fprintf(w, "user DateOfBirthday, %s\n", user.DateOfBirthday)
// 		fmt.Fprintf(w, "user PhoneNumber, %s\n", user.PhoneNumber)
// 		fmt.Fprintf(w, "user Login, %s\n", user.Login)
// 		fmt.Fprintf(w, "user PassHash, %s\n", user.PassHash)
// 		fmt.Fprintf(w, "ITS WOOOOOORK AAAIR %s", user.Name)

// 	} else {
// 		tmpl, _ := template.ParseFiles("templates/login.html")
// 		tmpl.Execute(w, nil)
// 	}
// }
