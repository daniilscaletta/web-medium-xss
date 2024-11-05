package main

import (
	"fmt"
	"html/template"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// type User struct {
// 	FirstName   string
// 	LastName    string
// 	DateOfBirth string
// 	PhoneNumber string
// 	Passport    string
// 	Login       string
// 	Password    string
// }

// var users = make(map[string]User)

// func registerHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		tmpl, _ := template.ParseFiles("register.html")
// 		tmpl.Execute(w, nil)
// 	} else if r.Method == http.MethodPost {
// 		user := User{
// 			FirstName:   r.FormValue("firstName"),
// 			LastName:    r.FormValue("lastName"),
// 			DateOfBirth: r.FormValue("dateOfBirth"),
// 			PhoneNumber: r.FormValue("phoneNumber"),
// 			Passport:    r.FormValue("passport"),
// 			Login:       r.FormValue("login"),
// 			Password:    r.FormValue("password"),
// 		}
// 		users[user.Login] = user
// 		fmt.Fprintf(w, "Регистрация успешна!")
// 	}
// }

// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		tmpl, _ := template.ParseFiles("login.html")
// 		tmpl.Execute(w, nil)
// 	} else if r.Method == http.MethodPost {
//login := r.FormValue("login")
// 	 			password := r.FormValue("password")

// 		user, exists := users[login]
// 		if !exists || user.Password != password {
// 			http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
// 			return
// 		}

// 		fmt.Fprintf(w, "Добро пожаловать, %s!", user.FirstName)
// 	}
// }

type User struct {
	Name        string
	Surname     string
	DateOfBirth string
	PhoneNumber int
	Email       string
	Passport    int
	Login       string
	Password    string
}

func home_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("home.html")
	tmpl.Execute(w, nil)
}

func register_page(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("register.html")
	tmpl.Execute(w, nil)
}

func login_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("login.html")
	tmpl.Execute(w, nil)
}

func handlerRequest() {
	http.HandleFunc("/home/", home_page)
	http.HandleFunc("/register/", register_page)
	http.HandleFunc("/login/", login_page)
	fmt.Println("Server is listening on http://localhost:1688")

	http.ListenAndServe(":1688", nil)
}

func main() {

	driver := "mysql"
	dsn := "root:qwerty@tcp(127.0.0.1:3306)/vkakids"

	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	fmt.Println("Connecting to the database...")

	insert, err := db.Prepare("INSERT INTO users(Name, Surname, DateOfBirthday, PhoneNumber, Email, Passport, Login, Password) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	insert.Exec("John", "Doe", "1990-01-01", 1234567890, "john@example.com", 1234123456, "johndoe", "qwerty123")

	defer insert.Close()

	fmt.Println("Connecting to the EEEENDDDD Table...")
	handlerRequest()

}
