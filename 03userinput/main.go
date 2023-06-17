package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our course:")

	// comma ok syntax || error ok
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Thanks for your input, ", input)
	fmt.Printf("Type of this ratings %T", input)

}