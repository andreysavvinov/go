package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 11; i += 2 {
			fmt.Println(i)
		}
	}()
	
	go func() {
		for i := 1; i < 9; i += 2 {
			fmt.Println(i)
		}
	}()

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("File 1 loaded")
		wg.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("File 2 loaded")
		wg.Done()
	}()

	go func(){
		time.Sleep(500 * time.Millisecond)
		fmt.Println("File 3 loaded")
		wg.Done()
	}()

	wg.Wait()
}
