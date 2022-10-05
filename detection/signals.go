package detection

import "obfuscation-detection/stats"

type Signals struct {
	StringLengthSummary       stats.SampleStatistics
	StringEntropySummary      stats.SampleStatistics
	CombinedStringEntropy     float64
	IdentifierLengthSummary   stats.SampleStatistics
	IdentifierEntropySummary  stats.SampleStatistics
	CombinedIdentifierEntropy float64
}
