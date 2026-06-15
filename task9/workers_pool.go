package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		//Имитация работы
		time.Sleep(1 * time.Second)
		results <- j * j * j
	}
}

func WorkersPool() {
	fmt.Println("-----WorkersPool-----")
	var Q int
	for {
		fmt.Println("Введите количество входных данных: ")
		_, err := fmt.Scanln(&Q)
		if err != nil || Q <= 0 {
			fmt.Println("Неправильный ввод. Введите положительное целое число:")
			var trash string
			fmt.Scanln(&trash)
			continue
		}
		break
	}
	input := make([]int, Q)
	for i := 0; i < Q; i++ {
		for {
			fmt.Println("Введите ", i+1, "-е число: ")
			_, err := fmt.Scanln(&input[i])
			if err != nil {
				fmt.Println("Неправильный ввод. Введите целое число:")
				var trash string
				fmt.Scanln(&trash)
				continue
			}
			break
		}
	}

	var wg sync.WaitGroup
	N := 3
	jobs := make(chan int, N)
	results := make(chan int, N)
	wg.Add(N)
	for i := 1; i <= N; i++ {
		go worker(i, jobs, results, &wg)
	}

	for _, j := range input {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	var out []int
	for r := range results {
		out = append(out, r)
	}
	fmt.Println("results: ", out)
}
