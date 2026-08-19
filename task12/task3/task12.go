package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	data, err := os.ReadFile("./task12/task3/users.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		fmt.Println("Ошибка разбора JSON:", err)
		return
	}

	f := excelize.NewFile()
	sheet := f.GetSheetName(0)
	f.SetCellValue(sheet, "A1", "Name")
	f.SetCellValue(sheet, "B1", "Age")

	for i, user := range users {
		row := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), user.Name)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), user.Age)
	}

	if err := f.SaveAs("./task12/task3/report.xlsx"); err != nil {
		fmt.Println("Ошибка сохранения:", err)
		return
	}
}