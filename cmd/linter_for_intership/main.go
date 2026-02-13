package linterforintership

import (
	logcheck "intershiplinter/log_check"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	analyzers := []*analysis.Analyzer{
		logcheck.LogEnglishLetterAnalyzer,
		logcheck.LogLowercaseAnalyzer,
		logcheck.LogSensitiveDataAnalyzer,
		logcheck.LogSpecialSymbolAnalyzer,
	}
	multichecker.Main(analyzers...)
}
