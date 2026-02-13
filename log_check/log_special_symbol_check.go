package logcheck

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var LogSpecialSymbolAnalyzer = &analysis.Analyzer{
	Name: "loglowercase",
	Doc:  "checks that log messages start with a lowercase letter",
	Run:  logSpecialSymbolCheck,
}

// Проверяет, что сообщения логов не содержат спецсимволов и эмодзи
func logSpecialSymbolCheck(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok || len(call.Args) == 0 || !isLogCall(pass, call) {
				return true
			}

			arg := call.Args[0]

			if containsSpecialOrEmojiExpr(arg) {
				pass.Reportf(arg.Pos(), "log message should not contain special symbols or emoji")
			}

			return true
		})
	}

	return nil, nil
}
