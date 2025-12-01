package main

import (
	"os"

	flag "github.com/spf13/pflag"
)

func parsing() *Config {
	var conf Config
	flag.BoolVarP(&conf.reverse, "reverse", "r", false, "Sorting data in reverse order")
	flag.BoolVarP(&conf.useDigit, "numbered", "n", false, "Interpret strings like a numbers")
	flag.BoolVarP(&conf.unique, "unique", "u", false, "Print only unique strings in output")
	flag.IntVarP(&conf.column, "column", "k", -1, "Compare using only k-column")
	flag.BoolVarP(&conf.sorted, "check", "c", false, "Check the data sorting")
	flag.Parse()
	return &conf
}

func main() {
	var sortMe Sortable

	sortMe.Config = parsing()

	sortMe.Fill(os.Stdin)

	sortMe.Sort()

	sortMe.Print(os.Stdout)
}
