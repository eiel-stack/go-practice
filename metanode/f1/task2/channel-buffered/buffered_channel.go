package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a buffered channel with a capacity of 10
	ch := make(chan int, 10)

	// Producer goroutine: sends 100 integers to the buffered channel
	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}

		// Close the channel after sending all values
		close(ch)
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	// Consumer goroutine: receives and prints values from the buffered channel
	go func() {
		// Signal completion when this goroutine finishes
		defer wg.Done()
		for v := range ch {
			fmt.Println("Received:", v)
		}
	}()

	// Wait for consumer goroutine to finish
	wg.Wait()
}
