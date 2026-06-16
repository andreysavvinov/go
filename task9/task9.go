package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// реализация первого пункта
	fmt.Println("----Channels----")
	chNumNotBuf := make(chan int)
	chNumBuf := make(chan int, 3)

	go func() {
		for i := 1; i < 10; i++ {
			num := rand.Int()
			fmt.Println("send notbuf ", num)
			chNumNotBuf <- num
			time.Sleep(1 * time.Second)
			fmt.Println("sent notbuf ", num)
		}
		close(chNumNotBuf)
	}()

	fmt.Println("Небуферизованный канал")
	fmt.Println("Пока данные не примутся горутина блокируется")
	for n := range chNumNotBuf {
		fmt.Println(n)
	}

	go func() {
		for i := 1; i < 10; i++ {
			num := rand.Int()
			fmt.Println("send buf ", num)
			chNumBuf <- num
			time.Sleep(1 * time.Second)
			fmt.Println("sent buf ", num)
		}
		close(chNumBuf)
	}()

	fmt.Println("Буферизованный канал")
	fmt.Println("Можем ждать пока буфер канала заполнится")
	for n := range chNumBuf {
		time.Sleep(5 * time.Second)
		fmt.Println(n)
	}
	//6 пункт
	WorkersPool()
	//2 пункт
	Select()
	//3 пункт
	LongOperation()
	//4 пункт
	AsyncCollect()
	//5 пункт
	Manualctxcancel()
	Ctxtimeout()
	//7 пункт
	CompareMutexAndAtomic()
	//8 пункт
	RaceCondition()
}
