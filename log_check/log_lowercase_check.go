package logcheck

import (
	"go/ast"
	"go/token"
	"strconv"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var LogLowercaseAnalyzer = &analysis.Analyzer{
	Name: "loglowercase",
	Doc:  "checks that log messages start with a lowercase letter",
	Run:  lowercaselogcheck,
}

// Проверяет, что все сообщения в логах начинаются с маленькой буквы
func lowercaselogcheck(pass *analysis.Pass) (interface{}, error) {
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

			literal, ok := call.Args[0].(*ast.BasicLit)
			if !ok || literal.Kind != token.STRING {
				return true
			}

			value := literal.Value
			// Если строка состоит только из ""
			if len(value) < 2 {
				return true
			}

			unquoted, err := strconv.Unquote(value)
			if err != nil || unquoted == "" {
				return true
			}

			r := []rune(unquoted)[0]
			if unicode.IsLetter(r) && !unicode.IsLower(r) {
				pass.Reportf(
					literal.Pos(),
					"log message should start with a lowercase letter",
				)
			}

			return true
		})
	}

	return nil, nil
}
