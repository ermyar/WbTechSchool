package main

import "fmt"

func removeInd[T any](ar []T, ind int) []T {
	if ind >= len(ar) || ind < 0 {
		// nothing to remove
		return ar
	}
	copy(ar[ind:], ar[(ind+1):])
	ar = ar[:len(ar)-1]
	return ar
}

func main() {
	ar := []any{123, "abs", 4, "wow", 7}
	// before remove
	fmt.Println(ar)

	// removing
	ar = removeInd(ar, 2)

	// final slice
	fmt.Println(ar, len(ar), cap(ar))
}
