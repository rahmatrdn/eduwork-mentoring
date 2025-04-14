package main

import (
	"log"
	"task-app/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"task-app/config"
	"task-app/routes"
)

func main() {
	// Load environment variables dari .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using default DB config.")
	}

	// Koneksi ke database
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Task{})


	// Inisialisasi Gin
	r := gin.Default()

	// Register semua route
	routes.SetupRoutes(r)

	// Jalankan server di port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
