package main

import "fmt"

type ExtractedStrings struct {
	rawLiterals []string
	strings     []string
}

type FindStringsMethod int

const (
	parsing FindStringsMethod = iota
	regex
)

func FindStrings(filePath string, method FindStringsMethod) (*ExtractedStrings, error) {
	switch method {
	case parsing:
		return FindStringsParsing(filePath)
	case regex:
		return FindStringsRegex(filePath)
	default:
		panic(fmt.Sprintf("FindStrings: unknown method %d", method))
	}
}
