package logcheck

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestLogEnglishLettersAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	analysistest.Run(t, testdata, LogEnglishLetterAnalyzer, "english_letter_ok", "english_letter_fail")
}
