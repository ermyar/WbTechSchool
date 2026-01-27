package main

import (
	"log/slog"
	"os"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"
)

const (
	inf int = 2147483646
)

func parsing() *Config {
	var conf Config
	var rng string
	flag.StringVarP(&conf.delimiter, "delimiter", "d", "\t", "Set custom delimiter")
	flag.BoolVarP(&conf.separated, "separated", "s", false, "Select only strings contained delimiter")
	flag.StringVarP(&rng, "fields", "f", "", "Shows only selected columns")
	flag.Parse()

	for _, s := range strings.Split(rng, `,`) {
		if len(s) == 0 {
			continue
		}
		interval := strings.Split(s, `-`)
		switch len(interval) {
		case 1:
			val, err := strconv.Atoi(interval[0])
			if err != nil {
				slog.Error("Error while parsing `fields` args", s, err)
				break
			}
			conf.ranges = append(conf.ranges, segment{val, val + 1})
		case 2:
			var seg segment = segment{-inf, inf}
			var err error

			if len(interval[0]) != 0 {
				seg.left, err = strconv.Atoi(interval[0])
				if err != nil {
					slog.Error("Error while parsing `fields` args", s, err)
					break
				}
			}

			if len(interval[1]) != 0 {
				seg.right, err = strconv.Atoi(interval[1])
				if err != nil {
					slog.Error("Error while parsing `fields` args", s, err)
					break
				}
			}
			seg.right++

			if seg.left <= seg.right {
				conf.ranges = append(conf.ranges, seg)
			}
		default:
			slog.Error("Wrong type of fields arg")
		}
	}

	return &conf
}

func init() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelError})))
}

func main() {
	var cut CutUtil

	cut.init(parsing())

	if err := cut.process(os.Stdin); err != nil {
		slog.Error("Error occured while handling", "err", err)
	}
}
