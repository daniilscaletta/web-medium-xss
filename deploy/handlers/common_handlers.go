package handlers

import (
	"example/v3/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())

	router.GET("/home/", HomePage)
	router.GET("/login/", middleware.RedirectIfAuthenticated(), LoginPage)
	router.POST("/login/", middleware.RedirectIfAuthenticated(), LoginHandler)
	router.GET("/register/", middleware.RedirectIfAuthenticated(), RegisterPage)
	router.POST("/register/", middleware.RedirectIfAuthenticated(), RegisterHandler)

	protected := router.Group("/")
	{
		protected.GET("/profile/", ProfilePage)
		protected.GET("/logout/", LogoutPage)
		protected.GET("/record/", RecordPage)
		protected.POST("/record/", RecordHandler)
	}

	return router
}
