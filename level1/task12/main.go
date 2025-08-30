package main

import "fmt"

func unique[T comparable](ar []T) []T {
	var ans []T

	mp := make(map[T]bool)
	for _, val := range ar {
		if _, exist := mp[val]; !exist {
			ans = append(ans, val)
			mp[val] = true
		}
	}

	return ans
}

func main() {
	fmt.Println(unique([]string{"cat", "cat", "dog", "cat", "tree"}))
}
