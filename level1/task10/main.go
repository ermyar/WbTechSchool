package main

import "fmt"

func groupThem(ar []float32, step int) map[int][]float32 {
	mp := make(map[int][]float32)
	for _, val := range ar {
		gr := int(val)
		gr = gr - (gr % step)
		mp[gr] = append(mp[gr], val)
	}
	return mp
}

func main() {
	ar := []float32{-25.4, -27.0, 13.0, 19.0,
		15.5, 24.5, -21.0, 32.5}
	mp := groupThem(ar, 10)
	fmt.Println(mp)
}
