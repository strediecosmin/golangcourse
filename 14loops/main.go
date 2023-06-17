package main

import "fmt"

func main()  {
	fmt.Println("Some loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for k,v := range days {
		fmt.Println("the key is: ", k , "The value is: ", v)
	}

	anyValue := 1

	for anyValue < 10 {
	
		if anyValue == 2 {
			goto lco
		}

		if anyValue == 5 {
			anyValue++
			continue
		}

		fmt.Println("Value is: ", anyValue)
		anyValue++
	}

	lco:
		fmt.Println("Thank you for learning you hit the value 2")

}