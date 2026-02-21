package lowercase

import (
	"errors"
	"go/ast"
	"strings"
	"unicode"
	"github.com/srKazuya/loglint/internal/engine"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "lowercase",
	Doc:  "checks if log message starts with a lowercase letter",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	engine.InspectLogs(pass, func(lit *ast.BasicLit, pkg, fn string) {
		val := strings.Trim(lit.Value, "`\"")
		if err := validate(val); err != nil {
			pass.Reportf(lit.Pos(), "logger %s: error in function %q: %s", pkg, fn, err)
		}
	})
	return nil, nil
}

func validate(s string) error {
	if len(s) == 0 { return nil }
	if unicode.IsUpper([]rune(s)[0]) {
		return errors.New("log message must begin with a lowercase letter")
	}
	return nil
}