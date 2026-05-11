package main 

import "fmt"

func main() {
	var intStack Stack[int]

	intStack.Push(10)
	intStack.Push(20)

	value, _ := intStack.Pop()
	fmt.Println(value)

	var stringStack Stack[string]

	stringStack.Push("Go")
	stringStack.Push("Stack")

	top, _ := stringStack.Peek()
	fmt.Println(top)
}