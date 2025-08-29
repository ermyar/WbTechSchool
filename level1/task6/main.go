package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}

	// cond variant
	{
		mu := sync.Mutex{}
		ok := false

		cond := sync.NewCond(&mu)

		for i := range 5 {
			wg.Add(1)
			go func() {
				cond.L.Lock()
				for !ok {
					cond.Wait()
				}
				cond.L.Unlock()
				fmt.Println("exitted from Cond goroutine", i)
				wg.Done()
			}()
		}

		time.Sleep(5 * time.Second)

		mu.Lock()
		ok = true
		mu.Unlock()

		cond.Broadcast()
	}
	wg.Wait()

	// Notify channel
	{
		ch := make(chan os.Signal, 5)
		signal.Notify(ch, os.Interrupt)

		fmt.Println("Please send us Interrupt signal (Ctrl+C for example)")
		wg.Add(1)
		go func() {
			<-ch
			fmt.Println("\nexitted because of signal")
			wg.Done()
		}()
	}
	wg.Wait()

	// context example
	{
		ctx, stop := context.WithTimeout(context.Background(), 5*time.Duration(time.Second))
		defer stop()

		wg.Add(1)
		go func(ctx context.Context) {
			defer wg.Done()

			cnt := 0
			for {
				select {
				case <-ctx.Done():
					fmt.Println("exitted because of done context")
					return
				default:
					fmt.Println("Do important work:", cnt)
					cnt++
				}
			}
		}(ctx)
	}
	wg.Wait()

	// runtime.Goexit example
	{
		wg.Add(1)
		go func() {
			defer wg.Done()

			fmt.Println("saying bye-bye to this goroutine")
			fmt.Println("finishing with Goexit")
			runtime.Goexit()
		}()
	}
	wg.Wait()

	// closing channel example
	{
		ch := make(chan interface{})

		// initiate workers
		for i := range 5 {
			wg.Add(1)
			go func() {
				defer wg.Done()

				for {
					val, exist := <-ch
					if !exist {
						fmt.Println("exitted worker", i)
						return
					}
					fmt.Println("accepted: ", val)
				}
			}()
		}

		// data producer
		for i := range 100 {
			ch <- i
		}
		close(ch)
	}
	wg.Wait()
}
