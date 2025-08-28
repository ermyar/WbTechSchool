package main

import (
	"fmt"
	"os"
	"strconv"
)

func readFromCh(ch <-chan interface{}) {
	for {
		val, exist := <-ch
		if !exist {
			break
		}
		fmt.Println("(Log) have read from ch: ", val)
	}
}

func initiateWorkers(ch chan interface{}, cnt int) {
	for range cnt {
		go readFromCh(ch)
	}
}
func main() {
	cnt, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	ch := make(chan interface{})

	initiateWorkers(ch, cnt)

	val := 0

	for {
		ch <- fmt.Sprintf("some cringe #%d", val)
		val++
		// time.Sleep(200 * time.Millisecond)
	}
}
