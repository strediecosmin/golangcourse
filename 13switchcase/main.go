package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println("Welcome to switch cases")

	// rand.Seed(time.Now().UnixNano())
	rand.NewSource(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Dice value is 2 and you can move 2 spots")
		fallthrough
	case 3:
		fmt.Println("Case 3")
		fallthrough
	default:
		fmt.Println("Default value is: ", diceNumber)
	}

}