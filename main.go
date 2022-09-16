package main

import (
	"fmt"
)

const malicious = "/home/maxfisher/obfuscation-detection/malicious.js"
const simple = "/home/maxfisher/obfuscation-detection/simple.js"
const filePath = malicious

// TODO
//  entropy of function names
//  entropy of variable names
//  string literal analysis
//  analysis of numeric arrays (entropy)

func main() {
	println("obfuscation-detection")

	extracted, err := FindStrings(filePath, parsing)

	if err == nil {
		errorType := "Error"
		if extracted != nil {
			errorType = "Non-fatal error"
		}
		fmt.Printf("%s occured while extracting strings: %s\n", errorType, err)
	}

	if extracted != nil {
		fmt.Printf("Found %d values: %v\n", len(extracted.strings), extracted.strings)
	} else {
		fmt.Println("Unable to extract strings")
	}
}
