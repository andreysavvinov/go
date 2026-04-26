package operators

import "fmt"

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Деление на ноль")
	}
	
	return a / b, nil
}