package logcheck

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var LogSensitiveDataAnalyzer = &analysis.Analyzer{
	Name: "logsensitive",
	Doc:  "checks that log messages contains sensitive values",
	Run:  sensitivecheck,
}

func sensitivecheck(pass *analysis.Pass) (interface{}, error) {
	if err := LoadConfig(); err != nil {
		return nil, err
	}

	if !config.EnableLowercaseRule {
		return nil, nil
	}

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if len(call.Args) == 0 {
				return true
			}

			if !isLogCall(pass, call) {
				return true
			}

			arg := call.Args[0]

			if containsSensitive(arg) {
				pass.Reportf(
					arg.Pos(),
					"log message should not contain sensitive data",
				)
			}

			return true
		})
	}

	return nil, nil
}
