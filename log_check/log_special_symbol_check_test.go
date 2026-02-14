package logcheck

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestLogSpecialSymbolAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	analysistest.Run(t, testdata, LogSpecialSymbolAnalyzer, "special_symbol_ok", "special_symbol_fail")
}
