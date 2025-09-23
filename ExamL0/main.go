package main

import "fmt"

func getMissing(ar []int) int {
	n := len(ar)
	sum := n * (n + 1) / 2

	for _, val := range ar {
		sum -= val
	}

	return sum
}

func main() {
	fmt.Println(getMissing([]int{3, 0, 1}))
	fmt.Println(getMissing([]int{1}))
	fmt.Println(getMissing([]int{0}))
	fmt.Println(getMissing([]int{0, 1}))
	fmt.Println(getMissing([]int{0, 2}))
}
