package main

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

type Config struct {
	after           int
	before          int
	caseInsensitive bool
	vInverted       bool
	fixed           bool
	counted         bool
	numbered        bool
	pattern         string
}

type GrepSearch struct {
	*Config
	regexp *regexp.Regexp
	count  int
}

func (gr *GrepSearch) Print(res result) {
	var sb strings.Builder

	if gr.counted {
		if res.useful {
			gr.count++
		}
		return
	}

	if len(res.file) > 0 {
		sb.WriteString(color.MagentaString(fmt.Sprintf("%s:", res.file)))
	}
	if gr.numbered {
		sb.WriteString(color.GreenString(fmt.Sprintf("%d:", res.num)))
	}
	sb.WriteString(res.line)
	fmt.Println(sb.String())
}

func (gr *GrepSearch) init(cfg *Config) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelError})))
	gr.Config = cfg

	if gr.fixed {
		return
	}

	regex := gr.pattern
	if gr.caseInsensitive {
		regex = "(?i)" + regex
	}
	re, err := regexp.Compile(regex)
	if err != nil {
		slog.Error("Wrong regexp pattern -- exitting")
		os.Exit(1)
	}
	gr.regexp = re
}

type result struct {
	file   string
	line   string
	num    int
	useful bool
}

func (gr *GrepSearch) Search(filename string, r io.Reader) ([]result, error) {
	var above []result
	var last = -gr.after - 1
	var res []result
	sc := bufio.NewScanner(r)

	for num := 1; sc.Scan(); num++ {
		str := sc.Text()

		if matches := gr.findMatches(str); gr.vInverted != (len(matches) != 0) {
			for _, i := range above {
				res = append(res, i)
			}
			res = append(res, result{file: filename,
				line: getColoredMatches(str, matches), num: num, useful: true})
			above = nil
			last = num
		} else if num <= last+gr.after {
			res = append(res, result{file: filename,
				line: str, num: num})
		} else if gr.before > 0 {
			above = append(above, result{file: filename, line: str, num: num})
		}

		if gr.before > 0 && num >= gr.before+1 && len(above) > gr.before {
			above = above[1:]
		}
	}
	return res, nil
}

func (gr *GrepSearch) SearchStdin(r io.Reader) error {
	var above []result
	var last = -gr.after - 1
	sc := bufio.NewScanner(r)
	for num := 1; sc.Scan(); num++ {
		str := sc.Text()

		if matches := gr.findMatches(str); gr.vInverted != (len(matches) != 0) {
			for _, i := range above {
				gr.Print(i)
			}
			gr.Print(
				result{line: getColoredMatches(str, matches), num: num, useful: true})
			above = nil
			last = num
		} else if num <= last+gr.after {
			gr.Print(
				result{line: str, num: num})
		} else if gr.before > 0 {
			above = append(above, result{line: str, num: num})
		}

		if gr.before > 0 && num >= gr.before+1 && len(above) > gr.before {
			above = above[1:]
		}
	}
	fmt.Println(gr.count)
	return nil
}

func (gr *GrepSearch) findMatches(str string) []string {
	if !gr.fixed {
		return gr.regexp.FindAllString(str, -1)
	}

	if strings.Contains(str, gr.pattern) {
		return []string{gr.pattern}
	}
	return nil
}

func getColoredMatches(s string, matches []string) string {
	for _, i := range matches {
		s = strings.ReplaceAll(s, i, color.RedString(i))
	}
	return s
}
