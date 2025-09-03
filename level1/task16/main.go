package main

import "fmt"

func quickSort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	pivot, left, right := ar[0], 0, len(ar)-1

	for {
		for ar[left] < pivot {
			left++
		}
		if right <= left {
			break
		}
		for ar[right] > pivot {
			right--
		}
		if right <= left {
			break
		}
		ar[left], ar[right] = ar[right], ar[left]
	}
	quickSort(ar[:left])
	quickSort(ar[left+1:])
}

func main() {
	ar := []int{58, 27, 39, 48, 18, 47, 90, 75, 69, 61}
	quickSort(ar)
	fmt.Println(ar)
}
