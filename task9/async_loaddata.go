package main

import (
	"fmt"
	"sync"
	"time"
)

func rout(id int, delay time.Duration) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		// имитация загрузки данных
		time.Sleep(delay)

		out <- fmt.Sprintf("worker %d finished after %v", id, delay)
	}()

	return out
}

func collect(channels ...<-chan string) <-chan string {
	out := make(chan string)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, ch := range channels {
		go func(c <-chan string) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func AsyncCollect() {
	fmt.Println("-----Async  collection- -----")
	ch1 := rout(1, 2*time.Second)
	ch2 := rout(2, 1*time.Second)
	ch3 := rout(3, 3*time.Second)

	start := time.Now()

	out := collect(ch1, ch2, ch3)

	for res := range out {
		fmt.Println(res)
	}

	fmt.Println("total:", time.Since(start))
}