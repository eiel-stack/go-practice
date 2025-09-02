package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var counter int64

	// Number of goroutines and increments per goroutine
	numGoroutines := 10
	incrementsPerGoroutine := 1000

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(incs int) {
			wg.Done()
			for j := 0; j < incs; j++ {
				// Perform atomic increment
				atomic.AddInt64(&counter, 1)
			}

		}(incrementsPerGoroutine)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final value of the counter
	fmt.Printf("Final count: %d\n", counter)

}
