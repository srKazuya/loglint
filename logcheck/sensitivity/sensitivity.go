package sensitivity

import (
	"errors"
	"go/ast"
	"strings"
	"github.com/srKazuya/loglint/logcheck/engine"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "nosecrets",
	Doc:  "prevents sensitive data leaking into logs",
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
	secrets := []string{"password", "secret", "token", "auth", "apiKey", "api_key"}
	lower := strings.ToLower(s)
	for _, key := range secrets {
		if strings.Contains(lower, key) {
			return errors.New("potential sensitive data leak")
		}
	}
	return nil
}