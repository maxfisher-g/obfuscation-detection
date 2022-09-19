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
	parsingFile FindStringsMethod = iota
	parsingString
	stringregexp
)

func readFile(filePath string) (string, error) {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(fileBytes), nil
}

func FindStrings(filePath string, method FindStringsMethod) (*ExtractedStrings, error) {
	// This case only exists for debug purposes, because apparently
	// parsing the file worked differently from parsing a string
	if method == parsingFile {
		return FindStringsParsingFile(filePath)
	}

	fileString, err := readFile(filePath)
	if err != nil {
		return nil, err
	}

	switch method {
	case parsingString:
		return FindStringsParsing(fileString)
	case stringregexp:
		return FindStringsRegexp(fileString)
	default:
		panic(fmt.Sprintf("unknown FindStringsMethod %d", method))
	}
}
