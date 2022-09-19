package main

import (
	"regexp"
	"strings"
)

// General reference for matching string literals
// https://blog.stevenlevithan.com/archives/match-quoted-string

// https://stackoverflow.com/a/10786066
var singleQuotedString = regexp.MustCompile(`'[^'\\]*(\\.[^'\\]*)*'`)
var doubleQuotedString = regexp.MustCompile(`"[^"\\]*(\\.[^"\\]*)*"`)
var backTickQuotedString = regexp.MustCompile("`[^`\\\\]*(\\\\.[^`\\\\]*)*`")

// https://stackoverflow.com/a/30737232
var singleQuotedString2 = regexp.MustCompile(`'(?:[^'\\]*(?:\\.)?)*'`)
var doubleQuotedString2 = regexp.MustCompile(`"(?:[^"\\]*(?:\\.)?)*"`)
var backTickQuotedString2 = regexp.MustCompile("`(?:[^`\\\\]*(?:\\\\.)?)*`")

func combineRegexp(regexps ...*regexp.Regexp) *regexp.Regexp {
	patterns := Map(regexps, func(r *regexp.Regexp) string { return r.String() })
	return regexp.MustCompile(strings.Join(patterns, "|"))
}

//goland:noinspection GoUnusedGlobalVariable
var anyQuotedString = combineRegexp(singleQuotedString, doubleQuotedString, backTickQuotedString)
var anyQuotedString2 = combineRegexp(singleQuotedString2, doubleQuotedString2, backTickQuotedString2)

func dequote(s string) string {
	if len(s) <= 2 {
		return ""
	} else {
		return s[1 : len(s)-1]
	}
}

func FindStringsRegexp(source string) (*ExtractedStrings, error) {
	allStrings := anyQuotedString2.FindAllString(source, -1)

	unquotedStrings := Map(allStrings, dequote)

	if allStrings != nil {
		return &ExtractedStrings{strings: unquotedStrings, rawLiterals: allStrings}, nil
	} else {
		return &ExtractedStrings{strings: []string{}, rawLiterals: []string{}}, nil
	}
}
