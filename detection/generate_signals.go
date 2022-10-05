package detection

import (
	"fmt"
	"obfuscation-detection/parsing"
	"obfuscation-detection/stats"
	"obfuscation-detection/stringentropy"
	"strings"
)

func getStrings(data *parsing.ParseResult) []string {
	var extractedStrings []string
	for _, d := range data.Literals {
		switch d.Type {
		case "string":
			extractedStrings = append(extractedStrings, d.Value.(string))
		case "float64":
		case "bool":
		default:
			// do nothing
		}
	}
	return extractedStrings
}

func getIdentifierNames(data *parsing.ParseResult) []string {
	identifierNames := make([]string, len(data.Identifiers))
	for _, ident := range data.Identifiers {
		identifierNames = append(identifierNames, ident.Name)
	}
	return identifierNames
}

// Returns
// Stats summary of string lengths
// Stats summary of string entropies
// Entropy of all symbols concatentated together
func characterAnalysis(symbols []string) (
	lengthSummary stats.SampleStatistics,
	entropySummary stats.SampleStatistics,
	combinedEntropy float64,
) {
	// measure character probabilities by looking at entire set of strings
	characterProbs := stringentropy.CharacterProbabilities(symbols)

	var entropies []float64
	var lengths []int
	for _, s := range symbols {
		entropies = append(entropies, stringentropy.CalculateEntropy(s, characterProbs))
		lengths = append(lengths, len(s))
	}

	lengthSummary = stats.CalculateSampleStatistics(lengths)
	entropySummary = stats.CalculateSampleStatistics(entropies)
	combinedEntropy = stringentropy.CalculateEntropy(strings.Join(symbols, ""), nil)
	return
}

// GenerateSignals
// Generates some data from parsing the given source file and prints it out
// Currently assumes the input is a valid JavaScript source file
//
// TODO Planned signals
//   - analysis of numeric arrays (entropy)
func GenerateSignals(jsSourceFile string) (*Signals, error) {
	data, err := parsing.ParseJS(jsSourceFile, false)
	if err != nil && data == nil {
		fmt.Printf("Error occured while reading %s: %v\n", jsSourceFile, err)
		return nil, err
	}

	signals := Signals{}

	stringLiterals := getStrings(data)
	identifierNames := getIdentifierNames(data)

	signals.StringLengthSummary, signals.StringEntropySummary, signals.CombinedStringEntropy =
		characterAnalysis(stringLiterals)

	signals.IdentifierLengthSummary, signals.IdentifierEntropySummary, signals.CombinedIdentifierEntropy =
		characterAnalysis(identifierNames)

	return &signals, nil
}
