package main

import (
	"fmt"
	"sync"
	//"sync/atomic"
)

var counter int
//var mu sync.Mutex

func RaceCondition() {
	fmt.Println("----RaceCondition----")
	var wg sync.WaitGroup

	// Запускаем 2 горутины
	for i := 0; i < 2; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < 100000; j++ {
				//mu.Lock()
				counter++ // <-- race condition
				//mu.Unlock
				//или
				//atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()

	fmt.Println("counter:", counter)
}