package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("./task12/task1/input.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	wordCount := make(map[string]int)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		line = strings.ToLower(line)

		words := strings.Fields(line)
		for _, word := range words {
			word = strings.Trim(word, ".,!?;:\"()[]{}")
			if word != "" {
				wordCount[word]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения:", err)
		return
	}

	outFile, err := os.Create("./task12/task1/output.csv")
	if err != nil {
		fmt.Println("Ошибка создания output.csv:", err)
		return
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	words := make([]string, 0, len(wordCount))
	for word := range wordCount {
		words = append(words, word)
	}
	sort.Strings(words)

	for _, word := range words {
		err := writer.Write([]string{
			word,
			fmt.Sprintf("%d", wordCount[word]),
		})
		if err != nil {
			fmt.Println("Ошибка записи:", err)
			return
		}
	}
}
