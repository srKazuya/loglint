package main

import (
	"github.com/srKazuya/loglint/logcheck"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(logcheck.AllAnalyzers()...)
}
