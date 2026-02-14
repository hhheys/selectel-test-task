package logcheck

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestLogLowercaseAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	analysistest.Run(t, testdata, LogLowercaseAnalyzer, "lowercase_ok", "lowercase_fail")
}
