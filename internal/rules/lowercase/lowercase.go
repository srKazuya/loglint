package lowercase

import (
	"errors"
	"go/ast"
	"strings"
	"unicode"

	"github.com/srKazuya/loglint/internal/config"
	"github.com/srKazuya/loglint/internal/engine"
	"golang.org/x/tools/go/analysis"
)

const RuleName = "lowercase"

var Analyzer = &analysis.Analyzer{
	Name: RuleName,
	Doc:  "checks if log message starts with a lowercase letter",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	if !config.IsEnabled(RuleName, true) {
		return nil, nil
	}

	engine.InspectLogs(pass, func(lit *ast.BasicLit, pkg, fn string) {
		val := strings.Trim(lit.Value, "`\"")
		severity := config.GetSeverity(RuleName, "error")
		if err := validate(val); err != nil {
			pass.Reportf(lit.Pos(),
				"[%s][%s] %s in %s.%s",
				severity,
				RuleName,
				err,
				pkg,
				fn,
			)
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