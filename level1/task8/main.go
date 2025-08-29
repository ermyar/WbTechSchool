package main

import "fmt"

// value in zero-indexation
func SetBit(input int64, index, value int) int64 {
	result := input
	var tmp int64 = (result ^ (1 << index))

	switch value {
	case 0:
		result &= tmp
	case 1:
		result |= tmp
	default:
		panic("wrong value for a bit")
	}
	return result
}

func main() {
	var val int64 = 5

	val = SetBit(val, 0, 0)

	fmt.Println(val) // 4
}
