package main

import (
	logcheck "github.com/Selectel/log-linter/log_check"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	analyzers := getAnalyzers()
	multichecker.Main(analyzers...)
}

// getAnalyzers возвращает список всех анализаторов
func getAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		logcheck.LogEnglishLetterAnalyzer,
		logcheck.LogLowercaseAnalyzer,
		logcheck.LogSensitiveDataAnalyzer,
		logcheck.LogSpecialSymbolAnalyzer,
	}
}

// New - интерфейс для golangci-lint plugin
func New(conf any) ([]*analysis.Analyzer, error) {
	return getAnalyzers(), nil
}
