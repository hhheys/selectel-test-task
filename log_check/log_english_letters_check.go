package logcheck

import (
	"go/ast"
	"go/token"
	"strconv"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var LogEnglishLetterAnalyzer = &analysis.Analyzer{
	Name: "logenglishletter",
	Doc:  "checks that log messages contain English letters",
	Run:  englishletterscheck,
}

// englishletterscheck проверяет, что все сообщения в логах начинаются с маленькой буквы
func englishletterscheck(pass *analysis.Pass) (interface{}, error) {
	if err := LoadConfig(); err != nil {
		return nil, err
	}

	if !config.EnableEnglishLettersRule {
		return nil, nil
	}

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok || len(call.Args) == 0 {
				return true
			}

			if !isLogCall(pass, call) {
				return true
			}

			literal, ok := call.Args[0].(*ast.BasicLit)
			if !ok || literal.Kind != token.STRING {
				return true
			}

			unquoted, err := strconv.Unquote(literal.Value)
			if err != nil || unquoted == "" {
				return true
			}

			for _, r := range unquoted {
				if unicode.IsLetter(r) && !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) {
					pass.Reportf(
						literal.Pos(),
						"log message should not contain non-English letters",
					)
					break // только одно сообщение на строку
				}
			}

			return true
		})
	}

	return nil, nil
}
