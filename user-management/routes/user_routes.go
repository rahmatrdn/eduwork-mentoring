package routes

import (
	"user-management/controllers"
	"user-management/middlewares"

	"github.com/gin-gonic/gin"
)
func UserRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected route group
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	auth.GET("/profile", controllers.GetProfile)
	auth.PUT("/profile", controllers.UpdateProfile)
	auth.PUT("/change-password", controllers.ChangePassword)
	r.POST("/forgot-password", controllers.ForgotPassword)
	r.POST("/reset-password", controllers.ResetPassword)


}