package main

import "fmt"

func producer(ch chan<- int) {
	for i := 0; i < 4; i++ {
		ch <- i // Send data to the channel
	}
	close(ch) // Close the channel to signal the end of data
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Println(num) // Print the received data
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	go consumer(ch)

	fmt.Scanln() // Wait for user input to prevent the program from exiting immediately
}
