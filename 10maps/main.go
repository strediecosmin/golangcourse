package main

import "fmt"

func main()  {
	fmt.Println("Welcome to maps series")

	myMap := map[string]int{
		"Cosmin": 31,
	}

	fmt.Println(myMap)

	secondMap := make(map[string]int)

	secondMap["Ana"] = 31
	secondMap["Bebeticulica"] = 21
	fmt.Println(secondMap)

	delete(secondMap, "Ana")
	fmt.Println("the new map of secondmap is: ", secondMap)

	// loop a map
	for k,v := range secondMap {
		fmt.Printf("The type for key is %T, and for type of value %T\n", k, v)
	}

	for _, v := range secondMap {
		fmt.Println("the value of the map is: ", v)
	}

}