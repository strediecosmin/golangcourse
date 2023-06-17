package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	createdDate := time.Date(2020, time.July, 02, 24, 59, 61, 66, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Mon"))

}
