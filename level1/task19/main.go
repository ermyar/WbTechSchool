package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// reversing string's bytes (replacing incorrect bytes order into Unicode replacement character)
func reverse(input string) string {
	bytes := []byte(input)
	var b strings.Builder
	for i := len(bytes); i > 0; {
		r, size := utf8.DecodeLastRune(bytes[:i])
		if r == utf8.RuneError {
			b.WriteRune(utf8.RuneError)
		} else {
			b.Write(bytes[i-size : i])
		}
		i -= size
	}
	return b.String()
}

func main() {
	// some test imitation
	for _, tc := range []struct {
		s string
		t string
	}{
		{"👩‍❤️‍💋‍👩", "👩‍💋‍️❤‍👩"},
		{"🏳️‍🌈", "🌈‍️🏳"},
		{"뢴", "ᆫᅬᄅ"},
		{"Привет", "тевирП"},
	} {
		if rev := reverse(tc.s); rev != tc.t {
			fmt.Println("Got wrong, actual", rev, "expected:", tc.t)
			return
		}
	}
	fmt.Println("Well, it works!")
}
