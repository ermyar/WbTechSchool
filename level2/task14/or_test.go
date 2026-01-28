package task14

import (
	"fmt"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	begin := time.Now()

	<-or(
		sig(2*time.Second),
		sig(time.Second),
		sig(time.Hour),
		sig(time.Minute),
		sig(5*time.Minute),
	)

	fmt.Println("Time passed:", time.Since(begin))
}
