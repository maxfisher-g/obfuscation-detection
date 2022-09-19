package main

import (
	"fmt"
	"os"
)

// TODO
//  entropy of function names
//  entropy of variable names
//  string literal analysis
//  analysis of numeric arrays (entropy)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <filename.js>\n", os.Args[0])
		return
	}

	filePath := os.Args[1]

	extractMethod := parsingFile
	extracted, err := FindStrings(filePath, extractMethod)

	if err != nil {
		errorType := "Error"
		if extracted != nil {
			errorType = "Non-fatal error"
		}
		fmt.Printf("%s occured while extracting strings: %s\n", errorType, err)
	}

	if extracted != nil {
		fmt.Printf("Found %d strings in: %s\n", len(extracted.strings), filePath)
		for _, s := range extracted.strings {
			println(s)
		}
	} else {
		fmt.Println("Unable to extract any strings from ", filePath)
	}
}
