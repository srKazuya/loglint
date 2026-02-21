package lowercase

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
	"unicode"
	"unicode/utf8"

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
		if lit.Kind != token.STRING {
			return
		}

		unquoted, err := strconv.Unquote(lit.Value)
		if err != nil {
			return
		}

		if err := validate(unquoted); err != nil {
			severity := config.GetSeverity(RuleName, "error")

			fix := buildFix(pass, lit, unquoted)

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
				SuggestedFixes: []analysis.SuggestedFix{fix},
			})
		}
	})

	return nil, nil
}

func buildFix(pass *analysis.Pass, lit *ast.BasicLit, unquoted string) analysis.SuggestedFix {
	if unquoted == "" {
		return analysis.SuggestedFix{}
	}

	firstRune, size := utf8.DecodeRuneInString(unquoted)
	lowerRune := unicode.ToLower(firstRune)

	file := pass.Fset.File(lit.Pos())
	startOffset := file.Offset(lit.Pos())

	editStart := file.Pos(startOffset + 1)
	editEnd := file.Pos(startOffset + 1 + size)

	return analysis.SuggestedFix{
		Message: "convert first letter to lowercase",
		TextEdits: []analysis.TextEdit{
			{
				Pos:     editStart,
				End:     editEnd,
				NewText: []byte(string(lowerRune)),
			},
		},
	}
}

func validate(s string) error {
	if s == "" {
		return nil
	}

	r, _ := utf8.DecodeRuneInString(s)
	if unicode.IsUpper(r) {
		return errors.New("log message must begin with a lowercase letter")
	}

	return nil
}