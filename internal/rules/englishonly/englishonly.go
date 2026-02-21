package englishonly

import (
	"errors"
	"go/ast"
	"strings"
	"unicode"

	"github.com/srKazuya/loglint/internal/config"
	"github.com/srKazuya/loglint/internal/engine"
	"golang.org/x/tools/go/analysis"
)

const RuleName = "englishonly"

var Analyzer = &analysis.Analyzer{
	Name: RuleName,
	Doc:  "checks that log messages use only the latin alphabet",
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
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.Is(unicode.Latin, r) {
			return errors.New("non-english characters detected (use latin alphabet only)")
		}
	}
	return nil
}