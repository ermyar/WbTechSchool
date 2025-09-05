package main

import "fmt"

func getIndLet(r rune) rune {
	if d := r - 'a'; 0 <= d && d < 26 {
		return d
	}
	if d := r - 'A'; 0 <= d && d < 26 {
		return d
	}
	return -1
}

// in case if we have only latin symbols
func checkUnique(s string) bool {
	var cnt [26]int
	for _, r := range s {
		if val := getIndLet(r); r != -1 {
			if cnt[val] == 1 {
				return false
			}
			cnt[val]++
		}
	}
	return true
}

func main() {
	fmt.Println("abcd", checkUnique("abcd"))
	fmt.Println("abCdefAaf", checkUnique("abCdefAaf"))
	fmt.Println("aabcd", checkUnique("aabcd"))
}
