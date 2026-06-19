package main

import "fmt"

type Role int 
const(
	Admin Role = iota
	Client
	Guest
)
//go:generate stringer -type=Role

func PrintStringerInterface(){
	fmt.Println(Guest)
}
