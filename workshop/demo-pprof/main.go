package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// https://github.com/gin-contrib/pprof
	"github.com/gin-contrib/pprof"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Register pprof routes under /debug/pprof
	pprof.Register(router)

	// Define a simple GET endpoint
	router.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Define an endpoint that simulates a memory leak
	router.GET("/memory-leak", func(c *gin.Context) {
		MemoryLeak()
		c.JSON(http.StatusOK, gin.H{
			"message": "Simulated memory leak",
		})
	})

	// Define an endpoint that simulates a CPU spike
	router.GET("/cpu-spike", func(c *gin.Context) {
		CPUSpike()
		c.JSON(http.StatusOK, gin.H{
			"message": "Simulated CPU spike",
		})
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	router.Run()
}

// cache to hold references to leaky slices
var cache [][]byte

func MemoryLeak() {
	// Simulate a memory leak by creating a large slice and not releasing it
	leakySlice := make([]byte, 10*1024*1024) // 10 MB
	// Add to cache to prevent garbage collection
	cache = append(cache, leakySlice)
}

func CPUSpike() {
	// Simulate a CPU spike by performing a large number of calculations
	for i := 0; i < 1e7; i++ {
		_ = i * i // Perform some calculation to consume CPU
	}
}
