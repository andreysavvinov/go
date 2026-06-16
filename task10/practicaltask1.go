package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

func ParsingUsersAndAgeChange(path string){
	var users []User

	data, err := os.ReadFile(path)

	if err != nil{
		fmt.Println("Error: ", err)
	}

	if err := json.Unmarshal(data, &users); err != nil{
		fmt.Println("Error: ", err)
	}

	for i := range users{
		users[i].Age += 1
	}
	
	fmt.Println(users)
}