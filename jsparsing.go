package main

import (
	"fmt"
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/parser"
)

func (extracted *ExtractedStrings) Enter(n ast.Node) (v ast.Visitor) {
	if parsedString, ok := n.(*ast.StringLiteral); ok {
		extracted.rawLiterals = append(extracted.rawLiterals, parsedString.Literal)
		extracted.strings = append(extracted.strings, parsedString.Value)
	}

	return extracted
}

func (extracted *ExtractedStrings) Exit(ast.Node) {}

func handleParsingResult(program *ast.Program, err error) (*ExtractedStrings, error) {
	if err != nil && program == nil {
		return nil, fmt.Errorf("parse error: %s", err)
	}

	e := ExtractedStrings{}
	ast.Walk(&e, program)

	return &e, err
}

func FindStringsParsing(fileString string) (*ExtractedStrings, error) {
	program, err := parser.ParseFile(nil, "", fileString, 0)
	return handleParsingResult(program, err)
}

func FindStringsParsingFile(filePath string) (*ExtractedStrings, error) {
	program, err := parser.ParseFile(nil, filePath, nil, 0)
	return handleParsingResult(program, err)
}
