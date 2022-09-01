package compiler

import (
	"unicode"
)

type Token struct {
	Type  Type
	Value string
}

type Type string

const (
	TypeParen  Type = "paren"
	TypeName   Type = "name"
	TypeNumber Type = "number"
	TypeString Type = "string"
)

func tokenizer(input string) []Token {
	current := 0
	tokens := []Token{}
	inputRunes := []rune(input)
	length := len(inputRunes)

	for current < length {
		char := inputRunes[current]

		if char == '(' {
			tokens = append(tokens, Token{Type: TypeParen, Value: "("})
			current++
			continue
		}

		if char == ')' {
			tokens = append(tokens, Token{Type: TypeParen, Value: ")"})
			current++
			continue
		}

		if char == '"' {
			value := ""
			current++
			char = inputRunes[current]

			for char != '"' {
				value += string(char)
				current++
				char = inputRunes[current]
			}

			current++
			char = inputRunes[current]

			tokens = append(tokens, Token{Type: TypeString, Value: value})

			continue
		}

		if unicode.IsSpace(char) {
			current++
			continue
		}

		if unicode.IsNumber(char) {
			value := ""

			for unicode.IsNumber(char) {
				value += string(char)
				current++
				char = inputRunes[current]
			}

			tokens = append(tokens, Token{Type: TypeNumber, Value: value})
		}

		if unicode.IsLetter(char) {
			value := ""

			for unicode.IsLetter(char) {
				value += string(char)
				current++
				char = inputRunes[current]
			}

			tokens = append(tokens, Token{Type: TypeName, Value: value})
		}
	}

	return tokens
}
