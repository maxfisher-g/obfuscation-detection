package main

import (
	"fmt"
	"os"
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
			println(s)
		}
	} else {
		fmt.Println("Unable to extract any strings from ", filePath)
	}

	if len(data.identifiers) > 0 {
		fmt.Printf("Found %d identifiers in: %s\n", len(data.identifiers), filePath)
		for _, ident := range data.identifiers {
			fmt.Printf("%s: %s\n", ident.Type, ident.Name)
		}
	} else {
		fmt.Println("Unable to extract any identifiers from ", filePath)
	}
}
