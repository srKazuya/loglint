package main

import (
	"github.com/srKazuya/loglint/internal"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(logcheck.AllAnalyzers()...)
}
