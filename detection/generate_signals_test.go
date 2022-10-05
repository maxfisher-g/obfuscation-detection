package detection

import (
	"math"
	"obfuscation-detection/stats"
	"obfuscation-detection/stringentropy"
	"testing"
)

const jsParserPath = "../parsing/babel-parser.js"

func singleStringEntropySummary(s string) stats.SampleStatistics {
	e := stringentropy.CalculateEntropy(s, nil)
	return stats.SampleStatistics{
		Size:      1,
		Mean:      e,
		Variance:  math.NaN(),
		Skewness:  math.NaN(),
		Quartiles: [5]float64{e, e, e, e, e},
	}
}

func singleStringLengthSummary(s string) stats.SampleStatistics {
	l := float64(len(s))
	return stats.SampleStatistics{
		Size:      1,
		Mean:      l,
		Variance:  math.NaN(),
		Skewness:  math.NaN(),
		Quartiles: [5]float64{l, l, l, l, l},
	}
}

func compareSummary(t *testing.T, name string, expected, actual stats.SampleStatistics) {
	if !expected.Equals(actual, 1e-4) {
		t.Errorf("%s summary did not match.\nExpected: %v\nActual: %v\n", name, expected, actual)
	}
}

func TestBasic(t *testing.T) {
	jsSource := `var a = "hello"`
	signals, err := GenerateSignals(jsParserPath, "", jsSource)
	if err != nil {
		t.Error(err)
	} else {
		expectedStringEntropySummary := singleStringEntropySummary("hello")
		expectedStringLengthSummary := singleStringLengthSummary("hello")
		expectedIdentifierEntropySummary := singleStringEntropySummary("a")
		expectedIdentifierLengthSummary := singleStringLengthSummary("a")

		compareSummary(t, "String literal entropy", expectedStringEntropySummary, signals.StringEntropySummary)
		compareSummary(t, "String literal lengths", expectedStringLengthSummary, signals.StringLengthSummary)
		compareSummary(t, "Identifier entropy", expectedIdentifierEntropySummary, signals.IdentifierEntropySummary)
		compareSummary(t, "Identifier lengths", expectedIdentifierLengthSummary, signals.IdentifierLengthSummary)

		// TODO compare combined entropies

	}
}
