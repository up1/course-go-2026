package main

import (
	"log"
	"os"

	"api/cache"
	"api/database"
	"api/product/handlers"
	"api/product/repositories"
	"api/product/services"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
	})
	defer rdb.Close()

	redisCache := cache.NewRedisCache(rdb)
	repo := repositories.NewProductRepository(db)
	svc := services.NewProductService(repo, redisCache)
	handler := handlers.NewProductHandler(svc)

	r := gin.Default()
	r.POST("/products", handler.Create)
	r.GET("/products/:id", handler.GetByID)

	port := getEnv("PORT", "8080")
	log.Printf("starting server on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
