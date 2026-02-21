package logcheck

import (
    "github.com/srKazuya/loglint/internal/rules/englishonly"
    "github.com/srKazuya/loglint/internal/rules/sensitivity"
    "github.com/srKazuya/loglint/internal/rules/lowercase" 
    "github.com/srKazuya/loglint/internal/rules/nosymbols" 
    
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