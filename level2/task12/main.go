package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
)

func parsing() *Config {
	var conf Config
	var closure int
	flag.IntVarP(&conf.after, "after", "A", 0, "Show \"num\" strings after")
	flag.IntVarP(&conf.before, "before", "B", 0, "Show \"num\" strings before")
	flag.IntVarP(&closure, "closure", "C", 0, "Show \"num\" strings in closure")
	flag.BoolVarP(&conf.numbered, "numbered", "n", false, "Prints number of the string")
	flag.BoolVarP(&conf.counted, "counted", "c", false, "Count matches with patterns")
	flag.BoolVarP(&conf.caseInsensitive, "case-insensitive", "i", false,
		"Search using case-intensive regex")
	flag.BoolVarP(&conf.vInverted, "inverted", "v", false, "Inverting searching")
	flag.BoolVarP(&conf.fixed, "fixed", "F", false, "Searching using fixed string (not a regex)")
	flag.Parse()

	conf.pattern = flag.Arg(0)
	conf.after = max(conf.after, closure)
	conf.before = max(conf.before, closure)
	return &conf
}

func (gr *GrepSearch) process() error {
	source := flag.Args()
	if len(source) == 0 {
		return fmt.Errorf("Nowhere to search, no pattern")
	}
	source = source[1:]
	if len(source) == 0 {
		return gr.SearchStdin(os.Stdin)
	}
	var answer []result

	for _, path := range source {
		m, err := filepath.Glob(path)
		if err != nil {
			slog.Error("wrong pattern", "err", err)
		}
		for _, file := range m {
			r, err := os.Open(file)
			if err != nil {
				slog.Error("Open error occured", "err", err)
			}
			res, err := gr.Search(file, r)
			if err != nil {
				slog.Error("Search error occured", "err", err)
			}
			answer = append(answer, res...)
			slog.Info("Finished handling", "file", file)
		}
	}
	for _, res := range answer {
		gr.Print(res)
	}
	if gr.counted {
		fmt.Println(gr.count)
	}
	return nil
}

func main() {
	var gr GrepSearch

	gr.init(parsing())

	if err := gr.process(); err != nil {
		slog.Error("Processed with mistakes", "err", err)
	}
}
