package main

import (
	"github.com/srKazuya/loglint/internal"
	"github.com/srKazuya/loglint/internal/config"
	"golang.org/x/tools/go/analysis/multichecker"

	_ "github.com/srKazuya/loglint/plugin"
)

func main() {
	config.Load("config.yaml")
	multichecker.Main(logcheck.AllAnalyzers()...)
}
