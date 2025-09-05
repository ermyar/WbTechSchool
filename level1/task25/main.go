package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	tt := time.Now()
	sleep(5 * time.Second)
	fmt.Println(time.Since(tt))
}
