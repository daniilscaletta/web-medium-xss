package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// func HandlerRequest() {
// 	http.HandleFunc("/home/", HomeHandler)
// 	http.HandleFunc("/register/", RegisterHandler)
// 	http.HandleFunc("/login/", LoginHandler)
// 	fmt.Println("Server is listening on http://localhost:1688")

// 	http.ListenAndServe(":1688", nil)
// }

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/home", HomeHandler)
	router.GET("/login", LoginPage)
	router.POST("/login", LoginHandler)
	router.GET("/register", RegisterPage)
	router.POST("/register", RegisterHandler)

	return router
}
