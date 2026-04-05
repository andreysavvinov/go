package main

import "fmt"

func main() {
	var hello string = "Hello world"
	fmt.Println(hello) // Hello world

	hello = "Hello Go"
	fmt.Println(hello) // Hello Go

	fmt.Println(hello) // Go Go Go Ole Ole Ole

	var numbers [5]int = [5]int{1,2,3,4,5}
    fmt.Print("type of  element: %T\n", numbers[0])     // 1
    fmt.Println(numbers[4])     // 5
    numbers[0] = 87
    fmt.Println(numbers[0])     // 87
}
