package logcheck

import (
    "github.com/srKazuya/loglint/internal/englishonly"
    "github.com/srKazuya/loglint/internal/sensitivity"
    "github.com/srKazuya/loglint/internal/lowercase" 
    "github.com/srKazuya/loglint/internal/nosymbols" 
    
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