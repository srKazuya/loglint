package sensitivity

import (
	"errors"
	"go/ast"
	"strings"

	"github.com/srKazuya/loglint/internal/config"
	"github.com/srKazuya/loglint/internal/engine"
	"golang.org/x/tools/go/analysis"
)

const RuleName = "sensivity"

var defaultPatterns = []string{"password", "secret", "auth", "apiKey", "api_key"}

var Analyzer = &analysis.Analyzer{
	Name: RuleName,
	Doc:  "prevents sensitive data leaking into logs",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	if !config.IsEnabled(RuleName, true) {
		return nil, nil
	}

	patterns := config.GetPatterns(RuleName, defaultPatterns)

	engine.InspectLogs(pass, func(lit *ast.BasicLit, pkg, fn string) {
		val := strings.Trim(lit.Value, "`\"")
		severity := config.GetSeverity(RuleName, "error")
		if err := validate(val, patterns); err != nil {
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

func validate(s string, patterns []string) error {
	lower := strings.ToLower(s)

	for _, key := range patterns {
		if strings.Contains(lower, strings.ToLower(key)) {
			return errors.New("potential sensitive data leak")
		}
	}

	return nil
}
