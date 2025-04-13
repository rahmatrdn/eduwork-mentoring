package main

import (
	"log"
	"net/http"
	"os"
	"product-management/config"
	"product-management/handler"
	"product-management/middleware"
	"product-management/repository"
	"product-management/service"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	db := config.InitDB()

	// Initialize dependencies
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// Setup router with middleware
	router := mux.NewRouter()
	router.Use(middleware.LoggingMiddleware)

	// Routes
	router.HandleFunc("/products", productHandler.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id}", productHandler.GetProduct).Methods("GET")
	router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")

	// Get port from env
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server berjalan pada port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
