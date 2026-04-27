package main

import (
	"fmt"

	. "github.com/andreysavvinov/go/task6/twice_slice"
)

func main() {
	//1
	a := 3
	b := 6.5
	sum := func(a, b float64) float64 { return a + b }
	fmt.Printf("Сумма двух чисел %f и %f:", a, b, sum)
	//2
	arr := [5]int{1, 2, 3, 4, 5}
	sli := arr[:3]
	fmt.Println("Удвоенный срез: ", SliceX2(sli))

}
