package logcheck

import (

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "addlint",
	Doc:  "reports integer additions",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	

	return nil, nil
}