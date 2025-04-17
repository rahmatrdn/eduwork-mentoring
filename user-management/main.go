package main

import (
	"user-management/config"
	"user-management/models"
	"user-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	// Panggil routing user
	routes.UserRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "User Management API ready 🚀"})
	})

	r.Run()
}
