package main

import (
	"fmt"
	"sync"
	"time"
)

type incrementor struct {
	mu  sync.Mutex
	cnt int64
}

var finish chan struct{}

func (i *incrementor) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-finish:
			return
		default:
			i.mu.Lock()
			i.cnt++
			i.mu.Unlock()
		}
	}
}

func main() {
	finish = make(chan struct{})
	ii := incrementor{cnt: 0}

	wg := sync.WaitGroup{}
	wg.Add(10)
	for range 10 {
		go ii.inc(&wg)
	}

	time.Sleep(5 * time.Second)
	close(finish)
	wg.Wait()

	fmt.Println(ii.cnt)
}
