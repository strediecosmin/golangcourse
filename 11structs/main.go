package main

import "fmt"

type repeatStructs struct {
	Name string
	Age int
	Country string
}

type differentStruct struct {
	name string
}

func main()  {
	fmt.Println("Welcome to structs lesson")

	newVariable := repeatStructs{
		"Cosmin", 31, "Romania",
	}
	fmt.Println(newVariable)

	anotheStruct := differentStruct{"Stredie"}

	fmt.Printf("just some details %+v\n", anotheStruct)
	fmt.Printf("The name is %+v\n", anotheStruct.name)
}