package main

import (
	"fmt"
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/file"
	"github.com/robertkrimen/otto/parser"
)

func (extracted *ExtractedStrings) Enter(n ast.Node) (v ast.Visitor) {
	if parsedString, ok := n.(*ast.StringLiteral); ok {
		extracted.rawLiterals = append(extracted.rawLiterals, parsedString.Literal)
		extracted.strings = append(extracted.strings, parsedString.Value)
	}

	return extracted
}

func (extracted *ExtractedStrings) Exit(n ast.Node) {}

func FindStringsParsing(filePath string) (*ExtractedStrings, error) {
	fileset := file.FileSet{}

	program, err := parser.ParseFile(&fileset, filePath, nil, 0)

	if err != nil && program == nil {
		return nil, fmt.Errorf("error parsing file: %s", err)
	}

	e := ExtractedStrings{}
	ast.Walk(&e, program)

	return &e, err
}
