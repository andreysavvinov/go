package main

import ("fmt"
"encoding/xml"
"os")

type Book struct{
	Title string `xml:"title"`
	Year int `xml:"year"`
}

type Library struct{
	Books []Book `xml:"book"`
}

func ParsingBooksAndYearChange(path string){
	var lib Library

	data, err := os.ReadFile(path)

	if err != nil{
		fmt.Println("Error: ", err)
	}

	if err := xml.Unmarshal(data, &lib); err != nil{
		fmt.Println("Error: ", err)
	}

	for i := range lib.Books{
		lib.Books[i].Year += 1
	}

	fmt.Println(lib)
}