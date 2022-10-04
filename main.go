package main

import (
	"fmt"
	"obfuscation-detection/parsing"
	"obfuscation-detection/stringentropy"
	"obfuscation-detection/utils"
	"os"
	"strings"
)

// TODO
//  entropy of identifier names
//  string literal analysis
//  analysis of numeric arrays (entropy)

func printProbabilityMap(m map[rune]float64) {
	mapStrings := utils.TransformMap(m, func(k rune, v float64) string {
		return fmt.Sprintf("%s: %.3f", string(k), v)
	})
	println(strings.Join(mapStrings, ", "))
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <filename.js>\n", os.Args[0])
		return
	}

	filePath := os.Args[1]

	//TestBabelParsing(filePath)
	//return

	data, err := parsing.ParseJS(filePath, false)
	if err != nil && data == nil {
		fmt.Printf("Error occured while extracting strings: %v\n", err)
		return
	}

	var e parsing.ExtractedStrings

	for _, d := range data.Literals {
		switch d.Type {
		case "string":
			e.Strings = append(e.Strings, d.Value.(string))
			e.RawLiterals = append(e.RawLiterals, d.RawValue)
			break
		case "float64":
		case "bool":
		default:
			// do nothing
		}
	}

	if len(e.Strings) > 0 {
		fmt.Printf("Found %d strings in: %s\n", len(e.Strings), filePath)
		for _, s := range e.Strings {
			entropy := stringentropy.CalculateEntropy(s, nil)
			entropyNormalised := stringentropy.CalculateNormalisedEntropy(s, nil)
			fmt.Printf("'%s' - entropy %.2f [%.1f%%]\n", s, entropy, 100*entropyNormalised)
		}
	} else {
		fmt.Println("Unable to extract any strings from ", filePath)
	}

	println()

	if len(data.Identifiers) > 0 {
		fmt.Printf("Found %d identifiers in: %s\n", len(data.Identifiers), filePath)
		identifierNames := make([]string, len(data.Identifiers))
		for _, ident := range data.Identifiers {
			identifierNames = append(identifierNames, ident.Name)
		}
		characterProbs := stringentropy.CharacterProbabilities(identifierNames)

		println("Character probabilities")
		printProbabilityMap(*characterProbs)
		println()

		for _, ident := range data.Identifiers {
			name := ident.Name
			dumbEntropy := stringentropy.CalculateEntropy(name, nil)
			dumbEntropyNormalised := stringentropy.CalculateNormalisedEntropy(name, nil)
			betterEntropy := stringentropy.CalculateEntropy(name, characterProbs)
			betterEntropyNormalised := stringentropy.CalculateNormalisedEntropy(name, characterProbs)
			fmt.Printf("%s: %s - naive entropy %.2f [%.1f%%], smart entropy %.2f [%.1f%%]\n",
				ident.Type, ident.Name, dumbEntropy, 100*dumbEntropyNormalised, betterEntropy, 100*betterEntropyNormalised)
		}

		combinedStrings := strings.Join(identifierNames, "")
		combinedEntropy := stringentropy.CalculateEntropy(combinedStrings, nil)
		combinedNormalisedEntropy := stringentropy.CalculateNormalisedEntropy(combinedStrings, nil)
		fmt.Printf("Combined entropy: %.2f [%.1f%%]\n", combinedEntropy, combinedNormalisedEntropy)

	} else {
		fmt.Println("Unable to extract any identifiers from ", filePath)
	}
}
