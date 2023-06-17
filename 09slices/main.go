package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Slices lesson repetition")

	var sliceLists = []string{"Mango","Apples"}
	fmt.Printf("The type is %T\n", sliceLists)

	sliceLists = append(sliceLists, "Mushroom")
	fmt.Println(sliceLists)

	sliceLists = sliceLists[:1]
	fmt.Println(sliceLists)

	highScores := make([]int, 4)

	highScores[0] = 222
	highScores[1] = 423
	highScores[2] = 224
	highScores[3] = 525
	// highScores[4] = 5555

	highScores = append(highScores, 555, 666, 777, 111)

	fmt.Println(highScores)
	
	sort.Ints(highScores)

	fmt.Println("sorted highscores: ", highScores)

	// remove a value from slices based on index

	courses := []string{"Golang", "React", "Java", "Ruby"}

	index := 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}