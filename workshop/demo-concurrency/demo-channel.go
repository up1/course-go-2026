package main

import "fmt"

type result struct {
	message string
}

func (r result) String() string {
	return "Result struct with message= " + r.message
}

func ping(ch chan<- result) {
	fmt.Println("Called ping function")
	ch <- result{message: "Pong"} // Signal that the work is done
}

func main() {

	// Create a channel to signal completion
	ch := make(chan result)

	fmt.Println("Process started")

	// Start a goroutine to perform some work
	go ping(ch)

	// Wait for the goroutine to finish
	result := <-ch
	fmt.Println(result)

	fmt.Println("Process done")

	// Wait for user input to prevent the program from exiting immediately
	fmt.Scanln()

}
