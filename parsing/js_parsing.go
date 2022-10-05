package parsing

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// BabelJSONElement
// Interfaces with output of babel-parser.js
type BabelJSONElement struct {
	SymbolType    string         `json:"type"`
	SymbolSubtype string         `json:"subtype"`
	Data          any            `json:"data"`
	Pos           [2]int         `json:"pos"`
	Array         bool           `json:"array"`
	Extra         map[string]any `json:"extra"`
}

func ParseJS(filePath string, printJson bool) (*ParseResult, error) {
	cmd := exec.Command("./parsing/babel-parser.js", filePath)
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
	} else {
		if printJson {
			println(jsonString)
		}
	}

	// convert the elements into more natural data structure
	result := ParseResult{}
	for _, element := range storage {
		switch element.SymbolType {
		case "Identifier":
			symbolSubtype := checkIdentifierType(element.SymbolSubtype)
			if symbolSubtype == other || symbolSubtype == unknown {
				break
			}
			result.Identifiers = append(result.Identifiers, ParsedIdentifier{
				Type: checkIdentifierType(element.SymbolSubtype),
				Name: element.Data.(string),
				Pos:  TextPosition{element.Pos[0], element.Pos[1]},
			})
		case "Literal":
			result.Literals = append(result.Literals, ParsedLiteral[any]{
				Type:     fmt.Sprintf("%T", element.Data),
				Value:    element.Data,
				RawValue: element.Extra["raw"].(string),
				InArray:  element.Array,
				Pos:      TextPosition{element.Pos[0], element.Pos[1]},
			})
		default:
			panic(fmt.Errorf("unknown element type for parsed symbol: %s", element.SymbolType))
		}
	}
	return &result, nil
}

func RunExampleParsing(filePath string) {
	parseResult, err := ParseJS(filePath, true)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		if ee, ok := err.(*exec.ExitError); ok {
			fmt.Printf("Process stderr:\n")
			fmt.Println(string(ee.Stderr))
		}
		return
	} else {
		fmt.Println("Completed without errors")
	}
	println()
	println("== Parsed Identifiers ==")
	for _, identifier := range parseResult.Identifiers {
		fmt.Println(identifier)
	}
	println()
	println("== Parsed Literals ==")
	for _, literal := range parseResult.Literals {
		fmt.Println(literal)
	}
}
