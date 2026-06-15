package main

import (
	"fmt"
	"time"
)

func goroutine1(ch1 chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		ch1 <- fmt.Sprintf("msg %d from goroutine1", i)
	}
}

func goroutine2(ch2 chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Second)
		ch2 <- fmt.Sprintf("msg %d from goroutine2", i)
	}
}

func Select() {
	fmt.Println("------Select-----")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goroutine1(ch1)
	go goroutine2(ch2)

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1, "accepted")

		case msg2 := <-ch2:
			fmt.Println(msg2, "accepted")

		case <-time.After(5 * time.Second):
			fmt.Println("Timeout")
			return
		}
	}
}