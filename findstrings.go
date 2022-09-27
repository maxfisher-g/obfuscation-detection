package main

import (
	"fmt"
	"os"
)

type ExtractedStrings struct {
	rawLiterals []string
	strings     []string
}

type FindStringsMethod int

const (
	parsing FindStringsMethod = iota
	stringregexp
)

func readFile(filePath string) (string, error) {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(fileBytes), nil
}

func findStringsParsing(filePath string) (*ExtractedStrings, error) {
	data, err := RunBabelParsing(filePath)
	if err != nil && data == nil {
		return nil, fmt.Errorf("parse error: %s", err)
	}

	e := ExtractedStrings{}
	for _, d := range data.literals {
		switch d.Type {
		case "string":
			e.strings = append(e.strings, d.Value.(string))
			e.rawLiterals = append(e.rawLiterals, d.RawValue)
		case "float64":
		case "bool":
		default:
			// do nothing
		}
	}

	return &e, err
}

func findStrings(filePath string, method FindStringsMethod) (*ExtractedStrings, error) {
	switch method {
	case parsing:
		return findStringsParsing(filePath)
	case stringregexp:
		fileString, err := readFile(filePath)
		if err != nil {
			return nil, err
		}
		return FindStringsRegexp(fileString)
	default:
		panic(fmt.Sprintf("unknown FindStringsMethod %d", method))
	}
}
