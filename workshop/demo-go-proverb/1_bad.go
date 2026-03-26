package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	data := make(map[string]int)

	fmt.Println("Starting")

	go func() {
		mu.Lock()
		data["key"] = 1
		fmt.Println("Data set in goroutine")
		mu.Unlock()
	}()

	// Prin data in main goroutine
	fmt.Println("Data received in main:", data["key"])

	// Alternatively, you can read the data in another goroutine
	go func() {
		mu.Lock()
		fmt.Println("Data received in goroutine:", data["key"])
		mu.Unlock()
	}()

	fmt.Println("Finished")

	// Press any key to exit
	fmt.Scanln()
}
