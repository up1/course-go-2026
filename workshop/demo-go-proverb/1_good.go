package main

import (
	"fmt"
)

func main() {

	fmt.Println("Starting")

	dataChan := make(chan map[string]int)

	// Set data in a goroutine
	go func() {
		dataChan <- map[string]int{"key": 1}
		fmt.Println("Data set in goroutine")
	}()

	data := <-dataChan
	fmt.Println("Data received in main:", data["key"])

	// Alternatively, you can read the data in another goroutine
	go func() {
		data := <-dataChan
		fmt.Println("Data received in goroutine:", data["key"])
	}()

	dataChan <- nil // Close the channel to signal completion

	fmt.Println("Finished")

	// Press any key to exit
	fmt.Scanln()
}
