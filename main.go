package main

import (
	"fmt"
	"os"
	"strings"
)

// TODO
//  entropy of identifier names
//  string literal analysis
//  analysis of numeric arrays (entropy)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <filename.js>\n", os.Args[0])
		return
	}

	filePath := os.Args[1]

	//TestBabelParsing()
	//return

	data, err := RunBabelParsing(filePath)
	if err != nil && data == nil {
		fmt.Printf("Error occured while extracting strings: %v\n", err)
		return
	}

	var e ExtractedStrings

	for _, d := range data.literals {
		switch d.Type {
		case "string":
			e.strings = append(e.strings, d.Value.(string))
			e.rawLiterals = append(e.rawLiterals, d.RawValue)
			break
		case "float64":
		case "bool":
		default:
			// do nothing
		}
	}

	if len(e.strings) > 0 {
		fmt.Printf("Found %d strings in: %s\n", len(e.strings), filePath)
		for _, s := range e.strings {
			entropy := StringEntropy(s, nil)
			entropyNormalised := StringEntropyNormalised(s, nil)
			fmt.Printf("'%s' - entropy %.2f [%.1f%%]\n", s, entropy, 100*entropyNormalised)
		}
	} else {
		fmt.Println("Unable to extract any strings from ", filePath)
	}

	println()

	if len(data.identifiers) > 0 {
		fmt.Printf("Found %d identifiers in: %s\n", len(data.identifiers), filePath)
		identifierNames := make([]string, len(data.identifiers))
		for _, ident := range data.identifiers {
			identifierNames = append(identifierNames, ident.Name)
		}
		characterProbs := CharacterProbabilities(identifierNames)

		println("Character probabilities")
		for ch, prob := range *characterProbs {
			fmt.Printf("%s: %.3f\n", string(ch), prob)
		}

		for _, ident := range data.identifiers {
			name := ident.Name
			dumbEntropy := StringEntropy(name, nil)
			dumbEntropyNormalised := StringEntropyNormalised(name, nil)
			betterEntropy := StringEntropy(name, characterProbs)
			betterEntropyNormalised := StringEntropyNormalised(name, characterProbs)
			fmt.Printf("%s: %s - naive entropy %.2f [%.1f%%], smart entropy %.2f [%.1f%%]\n",
				ident.Type, ident.Name, dumbEntropy, 100*dumbEntropyNormalised, betterEntropy, 100*betterEntropyNormalised)
		}

		combinedStrings := strings.Join(identifierNames, "")
		combinedEntropy := StringEntropy(combinedStrings, nil)
		combinedNormalisedEntropy := StringEntropyNormalised(combinedStrings, nil)
		fmt.Printf("Combined entropy: %.2f [%.1f%%]\n", combinedEntropy, combinedNormalisedEntropy)

	} else {
		fmt.Println("Unable to extract any identifiers from ", filePath)
	}
}
