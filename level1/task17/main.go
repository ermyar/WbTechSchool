package main

import "fmt"

func binSearch(ar []int, value int) int {
	left, right := 0, len(ar)
	var mid int
	for right-left > 1 {
		mid = (left + right) / 2
		if ar[mid] <= value {
			left = mid
		} else {
			right = mid
		}
	}
	if ar[left] == value {
		return left
	}
	return -1
}

func main() {
	fmt.Println(binSearch([]int{8, 18, 23, 31, 36, 38, 58, 67, 73, 76}, 30))
	fmt.Println(binSearch([]int{8, 18, 23, 31, 36, 38, 58, 67, 73, 76}, 38))
}
