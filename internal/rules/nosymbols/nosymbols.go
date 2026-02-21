package nosymbols

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

const RuleName = "nosymbols"

var Analyzer = &analysis.Analyzer{
	Name: RuleName,
	Doc:  "disallows emojis and special graphic symbols in log messages",
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

			fixed := removeSymbols(raw)

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
						Message: "remove symbols and emojis from log message",
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
		if unicode.IsSymbol(r) || unicode.Is(unicode.Sk, r) || unicode.Is(unicode.So, r) {
			return errors.New("special symbols or emojis are not allowed in logs")
		}
	}
	return nil
}

func removeSymbols(s string) string {
	var b strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) ||
			unicode.IsDigit(r) ||
			unicode.IsSpace(r) {
			b.WriteRune(r)
			continue
		}

		if unicode.IsPunct(r) {
			b.WriteRune(r)
		}
	}

	return b.String()
}