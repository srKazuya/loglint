package main

import (
	"github.com/srKazuya/loglint/internal"
	"golang.org/x/tools/go/analysis/multichecker"
	_ "github.com/srKazuya/loglint/plugin"
)

func main() {
	multichecker.Main(logcheck.AllAnalyzers()...)
}
