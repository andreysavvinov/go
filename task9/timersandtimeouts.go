package main

import (
	"fmt"
	"time"
)

func LongOperation() {
	fmt.Println("-----LongOperation------")
	fmt.Println("Долгая операция началась")
	timer := time.NewTimer(5 * time.Second)
	timeout := time.After(3 * time.Second)

	select {
	case <-timer.C:
		fmt.Println("Операция успешно завершена")
	case <-timeout:
		fmt.Println("Операция прервана по таймауту")	
		if !timer.Stop() {
			fmt.Println("Таймер уже сработал")
		}
	}
}