package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
)

func main()  {
	fmt.Println("Math in golang")

	// var firstNumber int = 2
	// var secondNumber float64 = 4

	// fmt.Println("The sum is: ", firstNumber + secondNumber)

	// Random Numbers
	// fmt.Println(rand.Intn(5) + 1)

	// Crypto Random number
	myRandomNumber, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		panic(err)
	}
	fmt.Println("My random number is: ", myRandomNumber)


}