package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type segment struct {
	left, right int
}

type Config struct {
	delimiter string
	ranges    []segment
	separated bool
}

type CutUtil struct {
	*Config
}

func (c *CutUtil) init(conf *Config) {
	c.Config = conf
	sort.Slice(c.ranges, func(i, j int) bool {
		return c.ranges[i].left < c.ranges[j].left || c.ranges[i].right < c.ranges[j].right
	})
}

func (c *CutUtil) process(r io.Reader) error {
	sc := bufio.NewScanner(r)

	for sc.Scan() {
		txt := sc.Text()
		c.handle(txt)
	}

	return nil
}

func (c *CutUtil) handle(s string) error {
	if c.separated && !strings.Contains(s, c.delimiter) {
		return nil
	}
	if len(c.ranges) == 0 {
		fmt.Println(s)
		return nil
	}
	ptr := 0
	var output []string
	for i, t := range strings.Split(s, c.delimiter) {
		for ptr < len(c.ranges) && i+1 >= c.ranges[ptr].right {
			ptr++
		}
		if ptr < len(c.ranges) && i+1 < c.ranges[ptr].left {
			continue
		}
		if ptr < len(c.ranges) && i+1 < c.ranges[ptr].right {
			output = append(output, t)
		}
	}
	if len(output) != 0 {
		fmt.Println(strings.Join(output, c.delimiter))
	}
	return nil
}
