package main

import "fmt"

type IdentifierType string

const (
	function IdentifierType = "Function"
	variable IdentifierType = "Variable"
)

func makeIdentifierType(s string) IdentifierType {
	switch s {
	case "Function":
		return function
	case "Variable":
		return variable
	default:
		panic(fmt.Errorf("unknown identifier type: %s", s))
	}
}

type ParsedIdentifier struct {
	Type IdentifierType
	Name string
	Pos  TextPosition
}

func (i ParsedIdentifier) String() string {
	return fmt.Sprintf("%s %s [pos %d:%d]", i.Type, i.Name, i.Pos.Row(), i.Pos.Col())
}

type ParsedLiteral[T any] struct {
	Type     string
	Value    T
	RawValue string
	InArray  bool
	Pos      TextPosition
}

func (l ParsedLiteral[T]) String() string {
	return fmt.Sprintf("%s %v (%s) [pos %d:%d]", l.Type, l.Value, l.RawValue, l.Pos.Row(), l.Pos.Col())
}

type BabelParseResult struct {
	identifiers []ParsedIdentifier
	literals    []ParsedLiteral[any]
}
