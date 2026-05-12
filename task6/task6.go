package main

import "fmt"

func SliceX2(sl []int) []int {
	sl2 := make([]int, len(sl))
	for i, el := range sl {
		sl2[i] = 2 * el
	}
	return sl2
}

func CountOfWords(mapa map[string]int)map[string]int {
	mapOfCounts := make(map[string]int)
	count := 0
	for name := range mapa {
		for word := range mapa {
			if name == word {
				count++
			}
		}
		mapOfCounts[name] = count
		count = 0
	}
	return mapOfCounts
}

type Person struct {
	Name string
	Salary  float32
}

func main() {
	//1
	a := 3.0
	b := 6.5
	sum := func(a, b float64) float64 { return a + b }
	fmt.Printf("Сумма двух чисел %.f и %.f: %.f\n", a, b, sum(a,b))
	//2
	arr := [5]int{1, 2, 3, 4, 5}
	sli := arr[:3]
	slc2 := SliceX2(sli)
	fmt.Println("Удвоенный срез: ", slc2)
	fmt.Println("Объединенный срез: ", append(sli, slc2...))
	//3
	mapa := map[string]int{"Алина": 18,
		"Алёна":     16,
		"Кристина":  16,
		"Мария":     21,
		"Анастасия": 23}
	mapa["Валерия"] = 24
	mapa["Валерия"] = 26
	delete(mapa, "Алёна")
	_, exists := mapa["Алёна"]
	countOfNames := CountOfWords(mapa)
	fmt.Printf("Количество каждого имени: %+v, конечно же\n", countOfNames)
	fmt.Println("Есть ли Алёна в списке? ", exists)
	//4
	chel := Person{Name: "Elon", Salary: 7850000000}
	fmt.Printf("Получится структура %+v\n", chel)
}
