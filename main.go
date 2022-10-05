package main

import (
	"fmt"
	"obfuscation-detection/parsing"
	"obfuscation-detection/stringentropy"
	"obfuscation-detection/utils"
	"os"
	"strings"
)

const defaultParserPath = "./parsing/babel-parser.js"

func getJsParserPath() (string, error) {
	customParserPath := os.Getenv("JS_PARSER")

	if len(customParserPath) > 0 {
		if _, err := os.Stat(customParserPath); err != nil {
			return "", fmt.Errorf("could not locate JS parser defined by environment variable (%v)", err)
		} else {
			return customParserPath, nil
		}
	} else {
		if _, err := os.Stat(defaultParserPath); err != nil {
			return "", fmt.Errorf("could not locate JS parser at default path %s (%v)", defaultParserPath, err)
		} else {
			return defaultParserPath, nil
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <filename.js>\n", os.Args[0])
		return
	}

	parserPath, err := getJsParserPath()

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	filePath := os.Args[1]
	printJson := true
	sourceString := "" // "var a = hello;"
	if len(sourceString) > 0 {
		runExampleAnalysis(parserPath, "", sourceString, printJson)
	} else {
		runExampleAnalysis(parserPath, filePath, "", printJson)
	}

}

func runExampleAnalysis(parserPath, jsFilePath, jsSource string, printJson bool) {
	data, err := parsing.ParseJS(parserPath, jsFilePath, jsSource, printJson)
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
		case "float64":
		case "bool":
		default:
			// do nothing
		}
	}

	if len(e.Strings) > 0 {
		fmt.Printf("Found %d strings in: %s\n", len(e.Strings), jsFilePath)
		for _, s := range e.Strings {
			entropy := stringentropy.CalculateEntropy(s, nil)
			entropyNormalised := stringentropy.CalculateNormalisedEntropy(s, nil)
			fmt.Printf("'%s' - entropy %.2f [%.1f%%]\n", s, entropy, 100*entropyNormalised)
		}
	} else {
		fmt.Println("Unable to extract any string literals")
	}

	println()

	if len(data.Identifiers) > 0 {
		fmt.Printf("Found %d identifiers in: %s\n", len(data.Identifiers), jsFilePath)
		identifierNames := make([]string, len(data.Identifiers))
		for _, ident := range data.Identifiers {
			identifierNames = append(identifierNames, ident.Name)
		}
		characterProbs := stringentropy.CharacterProbabilities(identifierNames)

		println("Character probabilities")
		utils.PrintProbabilityMap(*characterProbs)
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
		fmt.Println("Unable to extract any identifiers")
	}
}
