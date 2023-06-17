package main

import "fmt"

func main()  {
	fmt.Println("Welcome to arrays in golang")

	var fruitLists [3]string

	fruitLists[0] = "Apple"
	fruitLists[1] = "Tomato"

	fmt.Println(fruitLists)

	fmt.Println("The fruit list is: ", len(fruitLists))

	veggieList := [2]string{"carrots","potato"}

	fmt.Println("Veggie lists: ", veggieList)
	fmt.Printf("%T", veggieList)
}