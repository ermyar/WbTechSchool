package task11

import (
	"sort"
	"strings"
)

func getStat(s string) [33]int {
	ar := [33]int{}

	for _, r := range s {
		ar[r-'а']++
	}
	return ar
}

func GetAnagram(s ...string) map[string][]string {
	mp := make(map[[33]int][]int)
	for i := range s {
		s[i] = strings.ToLower(s[i])
		ar := getStat(s[i])
		mp[ar] = append(mp[ar], i)
	}
	ans := make(map[string][]string)
	for _, ar := range mp {
		if len(ar) == 1 {
			continue
		}
		for _, str := range ar {
			ans[s[ar[0]]] = append(ans[s[ar[0]]], s[str])
		}
	}
	for k := range ans {
		sort.Strings(ans[k])
	}
	return ans
}
