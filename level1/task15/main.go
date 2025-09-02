package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(len int) string {
	b := strings.Builder{}
	for range len {
		b.WriteString("世")
	}
	return b.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	// if we real want to get first 100 char, we'd use runes
	justString = string([]rune(v)[:100])
}

func main() {
	someFunc()
	fmt.Println(justString)
}
