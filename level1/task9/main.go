package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// write data to chan 1
	go func(n int) {
		defer close(ch1)
		for val := range n {
			ch1 <- val
		}
	}(1000)

	// handler
	go func() {
		defer close(ch2)
		for {
			val, exist := <-ch1
			if !exist {
				return
			}
			ch2 <- 2 * val
		}
	}()

	// final
	go func() {
		for {
			val, exist := <-ch2
			if !exist {
				return
			}
			fmt.Println("Final value:", val)
		}
	}()
}
