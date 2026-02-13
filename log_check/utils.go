package logcheck

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

// containsNonEnglishLetters проверяет, есть ли в строке буквы, не являющиеся латинскими
func containsNonEnglishLetters(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			if !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) {
				return true
			}
		} else {
			return true
		}
	}
	return false
}

// isLogCall проверяет, является ли вызов функцией логирования
func isLogCall(pass *analysis.Pass, call *ast.CallExpr) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	obj := pass.TypesInfo.Uses[sel.Sel]
	if obj == nil {
		return false
	}

	pkg := obj.Pkg()
	if pkg == nil {
		return false
	}

	// стандартный log
	if pkg.Path() == "log" {
		return true
	}

	// log/slog
	if pkg.Path() == "log/slog" {
		return true
	}

	if pkg.Path() == "go.uber.org/zap" {
		return true
	}

	return false
}

// containsSpecialOrEmoji проверяет, содержит ли строка специальные символы или эмодзи
func containsSpecialOrEmojiExpr(expr ast.Expr) bool {
	lit, ok := expr.(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return false
	}

	s, err := strconv.Unquote(lit.Value)
	if err != nil {
		return false
	}

	for _, r := range s {
		if !allowedChar(r) {
			return true
		}
	}

	return false
}

// allowedChar проверяет, является ли символ допустимым
func allowedChar(r rune) bool {
	if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
		return true
	}

	return false
}

// containsSensitive проверяет, содержит ли строка ключевые слова
func containsSensitive(expr ast.Expr) bool {
	switch v := expr.(type) {
	case *ast.BasicLit:
		if v.Kind != token.STRING {
			return false
		}
		unquoted, err := strconv.Unquote(v.Value)
		if err != nil {
			return false
		}
		return containsSensitiveKeyword(unquoted)

	case *ast.Ident:
		name := strings.ToLower(v.Name)
		return containsSensitiveKeyword(name)

	case *ast.BinaryExpr:
		if v.Op.String() == "+" {
			return containsSensitive(v.X) || containsSensitive(v.Y)
		}
	}

	return false
}

// containsSensitiveKeyword проверяет, содержит ли строка ключевое слово
func containsSensitiveKeyword(s string) bool {
	s = strings.ToLower(s)

	for _, k := range config.SensitiveKeywords {
		if s == k {
			return true
		}
	}

	return false
}
