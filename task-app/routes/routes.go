package routes

import (
	"github.com/gin-gonic/gin"

	"task-app/controllers"
	"task-app/middleware"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	// Auth routes
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	// Protected routes (butuh token JWT)
	protected := api.Group("/tasks")
	protected.Use(middleware.JWTAuth())

	protected.POST("/", controllers.CreateTask)
	protected.GET("/", controllers.GetTasks)
	protected.GET("/:id", controllers.GetTaskByID)
	protected.PUT("/:id", controllers.UpdateTask)
	protected.DELETE("/:id", controllers.DeleteTask)
	protected.PUT("/:id/complete", controllers.MarkTaskComplete)
}
