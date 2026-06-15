package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	goroutines = 100
	iterations = 100000
)

// Вариант с Mutex
func withMutex() int {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	start := time.Now()

	for i := 0; i < goroutines; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < iterations; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	fmt.Println("Mutex counter:", counter)
	fmt.Println("Mutex time:", time.Since(start))

	return counter
}

// Вариант с Atomic
func withAtomic() int64 {
	var counter int64
	var wg sync.WaitGroup

	start := time.Now()

	for i := 0; i < goroutines; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < iterations; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()

	fmt.Println("Atomic counter:", counter)
	fmt.Println("Atomic time:", time.Since(start))

	return counter
}

func CompareMutexAndAtomic() {
	fmt.Println("=== Mutex ===")
	withMutex()

	fmt.Println()

	fmt.Println("=== Atomic ===")
	withAtomic()
}