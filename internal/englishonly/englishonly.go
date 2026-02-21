package englishonly

import (
	"errors"
	"go/ast"
	"strings"
	"unicode"

	"github.com/srKazuya/loglint/internal/engine"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "englishonly",
	Doc:  "checks that log messages use only the latin alphabet",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	engine.InspectLogs(pass, func(lit *ast.BasicLit, pkg, fn string) {
		val := strings.Trim(lit.Value, "`\"")
		if err := validate(val); err != nil {
			pass.Reportf(lit.Pos(), "logger %s: function %q: %s", pkg, fn, err)
		}
	})
	return nil, nil
}

func validate(s string) error {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.Is(unicode.Latin, r) {
			return errors.New("non-english characters detected (use latin alphabet only)")
		}
	}
	return nil
}