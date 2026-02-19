package engine

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)


func InspectLogs(pass *analysis.Pass, check func(lit *ast.BasicLit, pkg, fn string)) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok { return true }

			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok { return true }

			obj := pass.TypesInfo.ObjectOf(sel.Sel)
			if obj == nil || obj.Pkg() == nil { return true }

			pkg := obj.Pkg().Path()
			if pkg == "log/slog" || pkg == "go.uber.org/zap" {
				if len(call.Args) > 0 {
					if lit, ok := call.Args[0].(*ast.BasicLit); ok && lit.Kind.String() == "STRING" {
						check(lit, pkg, sel.Sel.Name)
					}
				}
			}
			return true
		})
	}
}