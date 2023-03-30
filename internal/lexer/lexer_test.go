package lexer

import (
	"strings"
	"testing"
)

func Test_Tokenization(t *testing.T) {
	testCases := []struct {
		input          string
		expectedTokens []Token
	}{
		{
			input:          `"hello world"`,
			expectedTokens: []Token{{Type: TokenStringLiteral, ValueString: "hello world"}},
		},
		{
			input:          `"hello world""something""more text"`,
			expectedTokens: []Token{{Type: TokenStringLiteral, ValueString: "hello world"}, {Type: TokenStringLiteral, ValueString: "something"}, {Type: TokenStringLiteral, ValueString: "more text"}},
		},
		{
			input:          `"hello world"!{}`,
			expectedTokens: []Token{{Type: TokenStringLiteral, ValueString: "hello world"}, {Type: TokenBang}, {Type: TokenOpenBrace}, {Type: TokenCloseBrace}},
		},
		{
			input:          "\"hello world\" \"something\"	\n\"more text\"",
			expectedTokens: []Token{{Type: TokenStringLiteral, ValueString: "hello world"}, {Type: TokenStringLiteral, ValueString: "something"}, {Type: TokenStringLiteral, ValueString: "more text"}},
		},
	}

	for i, c := range testCases {
		lexer := New(strings.NewReader(c.input))

		for _, v := range c.expectedTokens {
			token, err := lexer.Read()

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if token.Type != v.Type {
				t.Errorf("case %d: Token.Type not as expected\nExpected: %s\nActual: %s", i, v.Type, token.Type)
			}

			if token.ValueString != v.ValueString {
				t.Errorf("case %d: Token.ValueString not as expected\nExpected: %s\nActual: %s", i, v.ValueString, token.ValueString)
			}

			if token.ValueInt != v.ValueInt {
				t.Errorf("case %d: Token.ValueInt not as expected\nExpected: %d\nActual: %d", i, v.ValueInt, token.ValueInt)
			}

			if token.ValueBool != v.ValueBool {
				t.Errorf("case %d: Token.ValueBool not as expected\nExpected: %v\nActual: %v", i, v.ValueBool, token.ValueBool)
			}
		}
	}
}

func Test_Peek(t *testing.T) {
	testCases := []struct {
		input           string
		expectedToken   Token
		inputTokenCount int
	}{
		{
			input:           `"hello world"`,
			expectedToken:   Token{Type: TokenStringLiteral, ValueString: "hello world"},
			inputTokenCount: 1,
		},
		{
			input:           `"hello world""something""more text"`,
			expectedToken:   Token{Type: TokenStringLiteral, ValueString: "hello world"},
			inputTokenCount: 2,
		},
		{
			input:           `"hello world"!{}`,
			expectedToken:   Token{Type: TokenStringLiteral, ValueString: "hello world"},
			inputTokenCount: 4,
		},
	}

	for i, c := range testCases {
		lexer := New(strings.NewReader(c.input))

		for range make([]byte, c.inputTokenCount) {
			token, err := lexer.Peek()

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if token.Type != c.expectedToken.Type {
				t.Errorf("case %d: Token.Type not as expected\nExpected: %s\nActual: %s", i, c.expectedToken.Type, token.Type)
			}

			if token.ValueString != c.expectedToken.ValueString {
				t.Errorf("case %d: Token.ValueString not as expected\nExpected: %s\nActual: %s", i, c.expectedToken.ValueString, token.ValueString)
			}

			if token.ValueInt != c.expectedToken.ValueInt {
				t.Errorf("case %d: Token.ValueInt not as expected\nExpected: %d\nActual: %d", i, c.expectedToken.ValueInt, token.ValueInt)
			}

			if token.ValueBool != c.expectedToken.ValueBool {
				t.Errorf("case %d: Token.ValueBool not as expected\nExpected: %v\nActual: %v", i, c.expectedToken.ValueBool, token.ValueBool)
			}
		}
	}
}
