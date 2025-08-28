package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	ch := make(chan interface{})

	go func() {
		for {
			val, exist := <-ch
			if !exist {
				return
			}
			fmt.Println("Accepted: ", val)
		}
	}()

	cnt := 0

	endCh := time.After(time.Duration(num) * time.Second)

	for {
		select {
		case <-endCh:
			close(ch)
			return
		default:
			ch <- cnt
			cnt++
		}
	}
}
