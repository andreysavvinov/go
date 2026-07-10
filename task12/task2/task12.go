package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	const fileName = "task12/task2/app.log"

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Ошибка создания файла:", err)
			return
		}
		file.Close()
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите текст (exit для выхода):")
	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		text := strings.TrimSpace(scanner.Text())
		if text == "exit" {
			fmt.Println("Работа завершена.")
			break
		}

		log.Println(text)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения:", err)
	}
}