package main

import (
	"fmt"
	"time"
)

func doWork(duration int) {
	fmt.Printf("Work started with %d ms\n", duration)
	time.Sleep(time.Duration(duration) * time.Millisecond)
	fmt.Printf("Work done with %d ms\n", duration)
}

func main() {
	// Start doWork in a new goroutine (fire-and-forget)
	go doWork(500)
	go doWork(200)
	go doWork(100)

	// Wait for a while to let the goroutine finish
	fmt.Scanln()
}
