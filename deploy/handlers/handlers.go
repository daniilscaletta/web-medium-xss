package handlers

import (
	"example/v3/models"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func HandlerRequest() {
	http.HandleFunc("/home/", homeHandler)
	http.HandleFunc("/register/", registerHandler)
	http.HandleFunc("/login/", loginHandler)
	fmt.Println("Server is listening on http://localhost:1688")

	http.ListenAndServe(":1688", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/home.html")
	tmpl.Execute(w, nil)

}
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		user := models.User{
			Name:           r.FormValue("Name"),
			Surname:        r.FormValue("Surname"),
			DateOfBirthday: r.FormValue("DateOfBirthday"),
			Email:          r.FormValue("Email"),
			PhoneNumber:    r.FormValue("PhoneNumber"),
			Passport:       r.FormValue("Passport"),
			Login:          r.FormValue("Login"),
			Password:       r.FormValue("Passhash"),
		}

		err := models.RegisterUser(&user)
		if err != nil {
			http.Error(w, "Ошибка при регистрации: "+err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		tmpl, _ := template.ParseFiles("templates/register.html")
		tmpl.Execute(w, nil)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		password := r.FormValue("password")

		user, err := models.AuthenticateUser(login, password)
		if err != nil {
			http.Error(w, "Ошибка авторизации: "+err.Error(), http.StatusUnauthorized)
			return
		}

		fmt.Fprintf(w, "Добро пожаловать, %s!", user.Name)
	} else {
		tmpl, _ := template.ParseFiles("templates/login.html")
		tmpl.Execute(w, nil)
	}
}
