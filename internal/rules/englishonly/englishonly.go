package englishonly

import (
	"errors"
	"fmt"
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
		raw := strings.Trim(lit.Value, "`\"")
		severity := config.GetSeverity(RuleName, "error")

		if err := validate(raw); err != nil {
			fixed := removeNonLatin(raw)

			quote := string(lit.Value[0])
			newValue := quote + fixed + quote

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
						Message: "remove non-latin characters",
						TextEdits: []analysis.TextEdit{
							{
								Pos:     lit.Pos(),
								End:     lit.End(),
								NewText: []byte(newValue),
							},
						},
					},
				},
			})
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

func removeNonLatin(s string) string {
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) {
			if unicode.Is(unicode.Latin, r) {
				b.WriteRune(r)
			}
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}