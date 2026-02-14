package logcheck

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestLogSensitiveDataAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	analysistest.Run(t, testdata, LogSensitiveDataAnalyzer, "sensitive_data_ok", "sensitive_data_fail")
}
