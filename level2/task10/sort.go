package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type Config struct {
	column   int
	reverse  bool
	useDigit bool
	unique   bool
	sorted   bool
}

type Pair struct {
	original string
	cmpStr   string
	cmpInt   float64
}

type Sortable struct {
	*Config
	a   []Pair
	cnt int
}

func (s *Sortable) Len() int {
	return len(s.a)
}

func (s *Sortable) Swap(i, j int) {
	s.a[i], s.a[j] = s.a[j], s.a[i]
	s.cnt++
}

func (s *Sortable) Less(i, j int) bool {
	if s.useDigit {
		if s.reverse {
			return s.a[i].cmpInt >= s.a[j].cmpInt
		}
		return s.a[i].cmpInt < s.a[j].cmpInt
	}

	if s.reverse {
		return s.a[i].cmpStr >= s.a[j].cmpStr
	}
	return s.a[i].cmpStr < s.a[j].cmpStr
}

func (s Sortable) Print(w io.Writer) {
	var lst string
	for i, pair := range s.a {
		if s.unique && i > 0 && lst == pair.original {
			continue
		}
		fmt.Fprintln(w, pair.original)
	}

	if s.sorted && s.cnt > 0 {
		fmt.Fprintln(w, "Data is not sorted!")
	}
}

func (s *Sortable) Fill(r io.Reader) {
	scan := bufio.NewScanner(r)
	for scan.Scan() {
		tmp := scan.Text()
		s.a = append(s.a, Pair{original: tmp})
	}
}

func (s *Sortable) Sort() {
	for i := range s.a {
		s.a[i].cmpStr = s.a[i].original
		if s.column >= 1 {
			if str := strings.Fields(s.a[i].original); len(str) >= s.column {
				s.a[i].cmpStr = str[s.column]
			}
		}
		if s.useDigit {
			if val, err := strconv.ParseFloat(s.a[i].cmpStr, 64); err == nil {
				s.a[i].cmpInt = float64(val)
			} else {
				fmt.Printf("Error while parsing %d to float64 with: %s\n", i+1, err)
			}
		}
	}
	sort.Sort(s)
}
