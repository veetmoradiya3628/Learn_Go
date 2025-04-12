package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{}
var wg sync.WaitGroup
var mut sync.Mutex // Mutex to protect the shared slice

func getStatusCode(endpoint string) {
	defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		mut.Lock() // Lock the mutex before modifying the shared slice
		signals = append(signals, endpoint)
		mut.Unlock()           // Unlock the mutex after modifying the shared slice
		defer res.Body.Close() // Close the response body when done
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
func main() {
	websitelist := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.github.com",
		"https://www.reddit.com",
		"https://www.stackoverflow.com",
		"https://www.quora.com",
		"https://www.medium.com",
		"https://www.wikipedia.org"}
	for _, site := range websitelist {
		go getStatusCode(site)
		wg.Add(1) // Increment the WaitGroup counter
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Signals:", signals)
	fmt.Println("All requests completed")
}
