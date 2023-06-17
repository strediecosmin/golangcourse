package main

import "fmt"

func main()  {
	defer fmt.Println("Deferred")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Defer keyword lesson")

	myDefer()

}

func myDefer()  {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}