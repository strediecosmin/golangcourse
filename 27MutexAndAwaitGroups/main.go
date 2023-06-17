package main

import (
	"fmt"
	"sync"
)

func main()  {
	fmt.Println("Race condition Lesson")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(4)
	go func (wg *sync.WaitGroup, m *sync.RWMutex)  {
		fmt.Println("First routine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func (wg *sync.WaitGroup, m *sync.RWMutex)  {
		fmt.Println("Second routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func (wg *sync.WaitGroup, m *sync.RWMutex)  {
		fmt.Println("Third routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func (wg *sync.WaitGroup, m *sync.RWMutex)  {
		fmt.Println("Third routine")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)
	
  wg.Wait()
	fmt.Println(score)

}

