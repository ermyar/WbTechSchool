package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	var sb strings.Builder

	// read-only mode (as i understand no copy and allocations)
	b := []byte(s)

	for lst, i := len(s), len(s)-1; i >= 0; i-- {
		if b[i] == ' ' {
			sb.Write(b[(i + 1):lst])
			sb.WriteRune(' ')
			lst = i
		} else if i == 0 {
			sb.Write(b[:lst])
		}
	}

	return sb.String()
}

func main() {
	str := "snow dog sun"
	trs := reverseWords(str)
	fmt.Println(str, len(str))
	fmt.Println(trs, len(trs))
}
