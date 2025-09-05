package main

import (
	"fmt"
	"sync"
)

func main() {
	sl := []int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}
	wg.Add(len(sl))

	squareIt := func(x int) {
		fmt.Println("x:", x, " square of x:", x*x)
		wg.Done()
	}

	for _, x := range sl {
		go squareIt(x)
	}
	wg.Wait()
}
