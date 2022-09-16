package main

import (
	"fmt"
	"os"
	"regexp"
)

// A good-looking reference page
// https://blog.stevenlevithan.com/archives/match-quoted-string

// https://stackoverflow.com/a/10786066
var singleQuotedString = regexp.MustCompile(`'([^'\\]*(\\.[^'\\]*)*)'`)
var doubleQuotedString = regexp.MustCompile(`"([^"\\]*(\\.[^"\\]*)*)"`)
var backTickQuotedString = regexp.MustCompile("`([^`\\\\]*(\\\\.[^`\\\\]*)*)`")

var anyQuotedString = regexp.MustCompile(fmt.Sprintf("%s|%s|%s",
	singleQuotedString.String(), doubleQuotedString.String(), backTickQuotedString.String()))

// https://stackoverflow.com/a/30737232
var doubleQuotedString2 = regexp.MustCompile(`"(?:[^"\\]*(?:\\.)?)*"`)

func FindStringsRegex(filePath string) (*ExtractedStrings, error) {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("FindStringsRegex(): error reading file - %s", err)
	}

	fileString := string(fileBytes)

	allStrings := anyQuotedString.FindAllString(fileString, -1)

	if allStrings != nil {
		return &ExtractedStrings{strings: allStrings, rawLiterals: []string{}}, nil
	} else {
		return &ExtractedStrings{strings: []string{}, rawLiterals: []string{}}, nil
	}
}
