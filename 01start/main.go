package main

import "fmt"

// export Public variable with CAPITAL LETTER
const LoginToken string = "eguhewuigiew"

func main()  {
	
	var username string = "Cosmin"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal float64 = 255.55555555555
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// default values and some aliases
	var someVariable int
	fmt.Println(someVariable)

	// implicit way to declare the variables

	website := "terraquantum.io"
	fmt.Println(website)

	fmt.Printf("The type of the variable is %T, and the value of it is %v", LoginToken, LoginToken)
	
}