package main

import (
	"fmt"

)

func main()  {
	a := 4
	squarePointer(&a)
	fmt.Println("value of a after calling the func: ", a)
	fmt.Printf("%p", initPerson())

}

func squarePointer(val *int)  {
	*val *= *val
	fmt.Println(*val, val)
}

type person struct {
	name string
	age int
}

func initPerson() *person {
	firstPerson := person{"Cosmin",31}
	fmt.Printf("value from inside the function to reflect heap %p \n", &firstPerson)
	return &firstPerson
}