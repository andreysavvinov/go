package int_input

import "fmt"

func Input() (a int) {
	fmt.Print("Введите целое значение: ")
	var err error
	for {
		_, err = fmt.Scanln(&a)
		if err != nil {
			fmt.Println("Ошибка: не int")

			var dummy string
			fmt.Scanln(&dummy)
			continue
		}
		return
	}
}
