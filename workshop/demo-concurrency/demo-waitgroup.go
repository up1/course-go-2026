package main

import (
	"net/http"
	"sync"
)

func fetchFromWebsite(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "Error fetching data"
	}
	defer resp.Body.Close()
	return "Data fetched successfully"
}

// Use channels to collect results from goroutines
func solution1(urls []string) {
	results := make(chan string)
	for _, url := range urls {
		go func(u string) {
			results <- fetchFromWebsite(u)
		}(url)
	}

	for range urls {
		result := <-results
		println(result)
	}
}

// Use wait groups to wait for all goroutines to finish
func solution2(urls []string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			result := fetchFromWebsite(u)
			println(result)
		}(url)
	}

	wg.Wait() // Wait for all goroutines to finish
}

func main() {
	urls := []string{
		"https://www.example.com",
		"https://www.google.com",
		"https://www.github.com",
	}

	solution1(urls)
	solution2(urls)

}
