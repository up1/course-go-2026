package main

import (
	"log"
	"os"

	"api/database"
	"api/product/handlers"
	"api/product/repositories"
	"api/product/services"

	"github.com/gin-gonic/gin"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "products.db"
	}

	db, err := database.NewDB(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	repo := repositories.NewProductRepository(db)
	service := services.NewProductService(repo)
	handler := handlers.NewProductHandler(service)

	r := gin.Default()
	r.POST("/products", handler.CreateProduct)
	r.GET("/products/:id", handler.GetProduct)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
