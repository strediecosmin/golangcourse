package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Pointers lesson")

	var pointer *int

	fmt.Println(pointer)

	randomNumber := 31

	var ptr = &randomNumber

	fmt.Println("Print value of reference to a pointer", ptr)
	fmt.Println("Print value of actual pointer", *ptr)
	fmt.Printf("The type of ptr is: %T", ptr)

	*ptr *= 2
	fmt.Println("The value after multiplying the pointer is: ", randomNumber)

	i, j := 42, 2701

	p := &i

	fmt.Println(*p)
	fmt.Println(j)

}
