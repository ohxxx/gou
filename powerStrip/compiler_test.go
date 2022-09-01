package compiler

import (
	"reflect"
	"testing"
)

var input = "(add 2 (subtract 4 2))"

var tokens = []Token{
	{Type: TypeParen, Value: "("},
	{Type: TypeName, Value: "add"},
	{Type: TypeNumber, Value: "2"},
	{Type: TypeParen, Value: "("},
	{Type: TypeName, Value: "subtract"},
	{Type: TypeNumber, Value: "4"},
	{Type: TypeNumber, Value: "2"},
	{Type: TypeParen, Value: ")"},
	{Type: TypeParen, Value: ")"},
}

func TestTokenizer(t *testing.T) {
	r := tokenizer(input)

	if !reflect.DeepEqual(r, tokens) {
		t.Error("\nExpected:", tokens, "\nGot:", r)
	}
}
