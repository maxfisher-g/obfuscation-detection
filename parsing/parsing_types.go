package parsing

import (
	"fmt"
)

type IdentifierType string

const (
	function  IdentifierType = "Function"
	variable  IdentifierType = "Variable"
	parameter IdentifierType = "Parameter"
	class     IdentifierType = "Class"
	other     IdentifierType = "Other"
	unknown   IdentifierType = "Unknown"
)

var allTypes = []IdentifierType{
	function,
	variable,
	parameter,
	class,
	other,
	unknown,
}

func checkIdentifierType(s string) IdentifierType {
	for _, typeName := range allTypes {
		if s == string(typeName) {
			return typeName
		}
	}
	return unknown
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

type ParseResult struct {
	Identifiers []ParsedIdentifier
	Literals    []ParsedLiteral[any]
}
