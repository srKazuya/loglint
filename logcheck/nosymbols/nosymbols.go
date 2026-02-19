package nosymbols

import (
	"errors"
	"go/ast"
	"strings"
	"unicode"

	"github.com/srKazuya/loglint/logcheck/engine"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "nosymbols",
	Doc:  "disallows emojis and special graphic symbols in log messages",
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
		if unicode.IsSymbol(r) {
			return errors.New("special symbols or emojis are not allowed in logs")
		}
	}
	return nil
}