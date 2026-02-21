package sensitivity

import (
	"errors"
	"fmt"
	"go/ast"
	"strings"

	"github.com/srKazuya/loglint/internal/config"
	"github.com/srKazuya/loglint/internal/engine"
	"golang.org/x/tools/go/analysis"
)

const RuleName = "sensivity"

var defaultPatterns = []string{
	"password",
	"secret",
	"auth",
	"apikey",
	"api_key",
}

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

		raw := strings.Trim(lit.Value, "`\"")
		severity := config.GetSeverity(RuleName, "error")

		if err := validate(raw, patterns); err != nil {

			pass.Report(analysis.Diagnostic{
				Pos:     lit.Pos(),
				End:     lit.End(),
				Message: fmt.Sprintf(
					"[%s][%s] %s in %s.%s",
					severity,
					RuleName,
					err,
					pkg,
					fn,
				),
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: "remove sensitive log message",
						TextEdits: []analysis.TextEdit{
							{
								Pos:     lit.Pos(),
								End:     lit.End(),
								NewText: []byte(`""`), 
							},
						},
					},
				},
			})
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