package main

import "fmt"

// first and second are sets.
func intersect[T comparable](first, second []T) []T {
	var ar []T
	for _, x := range first {
		for _, y := range second {
			if x == y {
				ar = append(ar, x)
				break
			}
		}
	}
	return ar
}

func main() {
	a := []int{6, 2, 3}
	b := []int{2, 5, 6}

	fmt.Println(intersect(a, b))
}
