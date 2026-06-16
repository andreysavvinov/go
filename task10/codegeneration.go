package main

import "fmt"

type Role int 
const(
	Admin Role = iota
	User
	Guest
)
//go:generate stringer -type=Role

func PrintStringerInterface(){
	fmt.Println(Guest)
}
