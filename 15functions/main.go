package main

import "fmt"

func newFunction()  {
	fmt.Println("call newFunction")
}

func adder(x,y int) int {
	return x + y
}

func proAdder(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	fmt.Printf("the type of values is: %T\n", values)
	return total
}

func main()  {
	fmt.Println("Functions in golang")
	newFunction()
	fmt.Println(adder(3,4))
	fmt.Println("call proAdder function result is: ", proAdder(1,2,3,4,5,6,6))

}