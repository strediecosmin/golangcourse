package main

import "fmt"

func main()  {
	fmt.Println("Control flow")

	loginCount := 10

	var result string

	if loginCount < 9 {
		result = "Just regular logins"
	} else if loginCount >= 10 {
		result = "It is greater or just equal"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	}

}