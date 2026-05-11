package main

import (
	"fmt"

	"github.com/andreysavvinov/go/StL/quick_sort"
	"github.com/andreysavvinov/go/StL/stack"
)

func main() {
	//stack testing
	var intStack stack.Stack[int]

	intStack.Push(10)
	intStack.Push(20)

	value, _ := intStack.Pop()
	fmt.Println(value)

	var stringStack stack.Stack[string]

	stringStack.Push("Go")
	stringStack.Push("Stack")

	top, _ := stringStack.Peek()
	fmt.Println(top)

	//quick_sort testing

	numbers := []int{5, 1, 9, 2, 7, 3}

	quick_sort.QuickSort(numbers, func(a, b int) bool {
		return a < b
	})

	fmt.Println(numbers)

	words := []string{"go", "cpp", "java", "python"}

	quick_sort.QuickSort(words, func(a, b string) bool {
		return a < b
	})

	fmt.Println(words)
}
