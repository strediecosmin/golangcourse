package main

import (
	"fmt"
	"sync"
)

func main()  {
	fmt.Println("Why do we need channels?")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}
	
	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)
	// receive only
	go func (ch <-chan int, wg *sync.WaitGroup)  {
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	// send only
	go func (ch chan<- int, wg *sync.WaitGroup)  {
		myCh <- 0
		close(myCh)
		// myCh <- 6
		wg.Done()		
	}(myCh, wg)

	wg.Wait()

	fmt.Printf("The type of waitgroups is: %T\n", wg)
	fmt.Printf("The value of waitgroup variable is: %#v", wg)
}