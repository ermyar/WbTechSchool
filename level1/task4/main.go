package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
)

func readFromCh(ch <-chan interface{}) {
	defer fmt.Println("exitted")

	for {
		val, exist := <-ch
		if !exist {
			return
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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	initiateWorkers(ch, cnt)

	val := 0

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stop producing to ch")
			close(ch)
			return
		default:
			ch <- fmt.Sprintf("some cringe #%d", val)
			val++
			// time.Sleep(200 * time.Millisecond)
		}
	}
}
