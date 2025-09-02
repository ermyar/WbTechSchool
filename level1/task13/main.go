package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// use rand values and int64
	a := rand.Int63()
	b := rand.Int63()

	fmt.Println("(Before) Variables and values: ", a, b)

	// swap
	b ^= a
	a ^= b
	b ^= a

	fmt.Println("(After) Variables and values: ", a, b)

}
