package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type BabelJSONElement struct {
	SymbolType    string         `json:"type"`
	SymbolSubtype string         `json:"subtype"`
	Data          any            `json:"data"`
	Pos           [2]int         `json:"pos"`
	Array         bool           `json:"array"`
	Extra         map[string]any `json:"extra"`
}

func RunBabelParsing(filePath string) (*BabelParseResult, error) {
	cmd := exec.Command("./babel-parser.js", filePath)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	jsonString := string(out)
	println("Decoding JSON")
	// parse JSON to get results as Go struct
	decoder := json.NewDecoder(strings.NewReader(jsonString))
	var storage []BabelJSONElement
	err = decoder.Decode(&storage)
	if err != nil {
		println("Failed on decoding the following JSON")
		println(jsonString)
		return nil, err
	}

	// convert the elements into more natural data structure
	result := BabelParseResult{}
	for _, element := range storage {
		switch element.SymbolType {
		case "Identifier":
			identifierType := checkIdentifierType(element.SymbolSubtype)
			if identifierType != unknown {
				result.identifiers = append(result.identifiers, ParsedIdentifier{
					Type: identifierType,
					Name: element.Data.(string),
					Pos:  TextPosition{element.Pos[0], element.Pos[1]},
				})
			}
			break
		case "Literal":
			result.literals = append(result.literals, ParsedLiteral[any]{
				Type:     fmt.Sprintf("%T", element.Data),
				Value:    element.Data,
				RawValue: element.Extra["raw"].(string),
				InArray:  element.Array,
				Pos:      TextPosition{element.Pos[0], element.Pos[1]},
			})
			break
		default:
			panic(fmt.Errorf("unknown element type for parsed symbol: %s", element.SymbolType))
		}
	}
	return &result, nil
}

func TestBabelParsing() {
	parseResult, err := RunBabelParsing("./test-strings.js")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		if ee, ok := err.(*exec.ExitError); ok {
			fmt.Printf("Process stderr:\n")
			fmt.Println(string(ee.Stderr))
		}
	} else {
		fmt.Println("Completed without errors")
	}
	println()
	println("== Parsed Identifiers ==")
	for _, identifier := range parseResult.identifiers {
		fmt.Println(identifier)
	}
	println()
	println("== Parsed Literals ==")
	for _, literal := range parseResult.literals {
		fmt.Println(literal)
	}
}
