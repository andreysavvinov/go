package main

import (
	"errors"
	"fmt"
)

type Printable interface {
	Print()
}

type Employee struct {
	Name     string
	Age      int
	Position string
	Salary   float32
}

var shtat []Employee

var position_salary map[string]float32

func (e *Employee) Print() {
	fmt.Println(e.Info())
}

func (e *Employee) Info() string {
	return fmt.Sprintf(
		"Name: %s, Age: %d, Position: %s, Salary: %.2f",
		e.Name,
		e.Age,
		e.Position,
		e.Salary,
	)
}

func FindEmployee(name string) (Employee, error) {
	for _, emp := range shtat {
		if emp.Name == name {
			return emp, nil
		}
	}

	return Employee{}, errors.New("employee not found")
}

func AverageAge() float64 {
	if len(shtat) == 0 {
		return 0
	}

	sum := 0

	for _, emp := range shtat {
		sum += emp.Age
	}

	return float64(sum) / float64(len(shtat))
}

func AddEmployee(p Employee) {
	shtat = append(shtat, p)
}

func UpdatePSmap() {
	position_salary = make(map[string]float32)
	countMap := make(map[string]float32)
	sumMap := make(map[string]float32)

	for _, p := range shtat {
		countMap[p.Position]++
		sumMap[p.Position] += p.Salary
	}

	for position, sum := range sumMap {
		position_salary[position] = sum / countMap[position]
	}
}

func main() {

	e1 := Employee{
		Name:     "Ivan",
		Age:      25,
		Position: "Developer",
		Salary:   1500,
	}

	e2 := Employee{
		Name:     "Anna",
		Age:      30,
		Position: "Designer",
		Salary:   1200,
	}

	e3 := Employee{
		Name:     "Petr",
		Age:      35,
		Position: "Developer",
		Salary:   2000,
	}

	shtat = append(shtat, e1, e2, e3)

	fmt.Println("=== Employees ===")

	for _, emp := range shtat {
		fmt.Println(emp.Info())
	}

	fmt.Println("\n=== Search ===")

	emp, err := FindEmployee("Anna")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(emp.Info())
	}

	// Средний возраст
	fmt.Println("\n=== Average Age ===")
	fmt.Printf("Average age: %.2f\n", AverageAge())

	UpdatePSmap()

	fmt.Println("\n=== Average Salary By Position ===")

	for position, salary := range position_salary {
		fmt.Printf("%s : %.2f\n", position, salary)
	}

}
