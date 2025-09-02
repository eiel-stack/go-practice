package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create an unbuffered channel to communicate integers between goroutines
	ch := make(chan int)

	// Use WaitGroup to wait for consumer goroutine to finish
	var wg sync.WaitGroup
	wg.Add(1)

	// Producer goroutine : sends integers from 1 to 10 to the channel
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		// Close the channel after sending all values
		close(ch)
	}()

	// Consumer goroutine : receives integers from the channel and prints them
	go func() {
		// Signal completion when this goroutine finishes
		defer wg.Done()
		for i := range ch {
			fmt.Println("Received:", i)
		}
	}()

	// Wait for consumer goroutine to finish
	wg.Wait()
}
