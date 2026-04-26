package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/andreysavvinov/go/task5/int_input"
	"github.com/andreysavvinov/go/task5/operators"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
div_cycle:
	for {
		a := int_input.Input()
		b := int_input.Input()
		res, err := operators.Divide(a, b)

		if err != nil {
			fmt.Println("Возникла ошибка при делении: ", err)
		} else {
			fmt.Println("Результат деления: ", res)
			switch {
			case res > 1 && res < 10:
				fmt.Println("Результат средний")
			case res > 10:
				fmt.Println("Результат большой")
			default:
				fmt.Println("Результат маленький или ноль")
			}
		}

		for {
			fmt.Println("Хотите повторить (вводить новые числа для деления)? y - да, n - нет")

			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			switch input {
			case "y", "Y":
				continue div_cycle
			case "n", "N":
				break div_cycle
			default:
				fmt.Println("Введите y или n")
			}
		}

	}
}
