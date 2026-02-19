package logcheck

import (
    "github.com/srKazuya/loglint/logcheck/englishonly"
    "github.com/srKazuya/loglint/logcheck/sensitivity"
    "github.com/srKazuya/loglint/logcheck/lowercase" 
    "github.com/srKazuya/loglint/logcheck/nosymbols" 
    
    "golang.org/x/tools/go/analysis"
)

func AllAnalyzers() []*analysis.Analyzer {
    return []*analysis.Analyzer{
        englishonly.Analyzer,
        sensitivity.Analyzer,
        lowercase.Analyzer,
        nosymbols.Analyzer,
    }
}