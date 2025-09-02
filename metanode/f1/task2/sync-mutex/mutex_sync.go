package main

import (
	"fmt"
	"sync"
)

// Counter holds a count and a mutex for synchronization
type Counter struct {
	mu    sync.Mutex
	count int
}

func main() {

	counter := Counter{}
	var wg sync.WaitGroup

	// Number of goroutines and increments per goroutine
	numGoroutines := 10
	incrementsPerGoroutine := 1000

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(incs int) {
			defer wg.Done()
			for j := 0; j < incs; j++ {
				// Lock the mutex to protect the counter
				counter.mu.Lock()
				counter.count++
				// Unlock the mutex after modifying the counter
				counter.mu.Unlock()
			}
		}(incrementsPerGoroutine)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final value of the counter
	fmt.Printf("Final count: %d\n", counter.count)
}
