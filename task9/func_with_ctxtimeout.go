package main

import (
	"context"
	"fmt"
	"time"
)

func Ctxtimeout() {
	fmt.Println("-----Пакет context-----")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res := make(chan string)
	go func() {
		fmt.Println("operation started")
		time.Sleep(4 * time.Second)
		res <- "operation end"
	}()
	select {
	case <-ctx.Done():
		fmt.Println("ctx timeout: ", ctx.Err())
	case result := <-res:
		fmt.Println(result)
	}
}

func Manualctxcancel() {
ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	work := func() {
		t := time.NewTicker(200 * time.Millisecond)
		defer t.Stop()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Manual cancel:", ctx.Err())
				return

			case <-t.C:
				fmt.Println("tick")
			}
		}
	}

	go work() 

	time.Sleep(3 * time.Second)

	fmt.Println("calling cancel()")
	cancel()

	time.Sleep(500 * time.Millisecond)
}
