package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Create REST API with net/http package

	http.HandleFunc("/get-all", process)
	// Print a message to indicate the server is running
	fmt.Println("Server is running on http://localhost:8080/get-all")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func process(w http.ResponseWriter, r *http.Request) {
	// Create channels to receive data from workers
	profileCh := make(chan string)
	ordersCh := make(chan string)

	// Start workers to fetch data concurrently
	go getProfile(profileCh)
	go getOrders(ordersCh)

	// Wait and combine results
	data := fmt.Sprintf("Result ::  %s, %s", <-profileCh, <-ordersCh)
	fmt.Println("Combined Result:", data)

	// Send the combined result back to the client
	w.Write([]byte(data))
}

func getProfile(ch chan<- string) {
	// Random delay(100-500ms) to simulate fetching data
	time.Sleep(time.Duration(100+rand.Intn(400)) * time.Millisecond)

	fmt.Println("Fetching user profile...")
	ch <- "User: Somkiat"
}

func getOrders(ch chan<- string) {
	// Random delay(100-500ms) to simulate fetching data
	time.Sleep(time.Duration(100+rand.Intn(400)) * time.Millisecond)
	fmt.Println("Fetching recent orders...")
	ch <- "Orders: 5 Recent"
}
