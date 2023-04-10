package lexer

import (
	"io"
	"strings"
	"testing"
)

func Test_Tokenization(t *testing.T) {
	testCases := []struct {
		input          io.Reader
		expectedTokens []Token
	}{
		{
			input:          strings.NewReader(`"hello world"`),
			expectedTokens: []Token{{Type: TokenStringLiteral, ValueString: "hello world"}},
		},
		{
			input:          strings.NewReader(`"hello world""something""more text"`),
			expectedTokens: []Token{{Type: TokenStringLiteral, ValueString: "hello world"}, {Type: TokenStringLiteral, ValueString: "something"}, {Type: TokenStringLiteral, ValueString: "more text"}},
		},
		{
			input:          strings.NewReader(`"hello world"!{}`),
			expectedTokens: []Token{{Type: TokenStringLiteral, ValueString: "hello world"}, {Type: TokenBang}, {Type: TokenOpenBrace}, {Type: TokenCloseBrace}},
		},
		{
			input:          strings.NewReader("\"hello world\" \"something\"	\n\"more text\""),
			expectedTokens: []Token{{Type: TokenStringLiteral, ValueString: "hello world"}, {Type: TokenStringLiteral, ValueString: "something"}, {Type: TokenStringLiteral, ValueString: "more text"}},
		},
		{
			input:          strings.NewReader(`123`),
			expectedTokens: []Token{{Type: TokenIntLiteral, ValueInt: 123}},
		},
		{
			input:          strings.NewReader(`123 456 99`),
			expectedTokens: []Token{{Type: TokenIntLiteral, ValueInt: 123}, {Type: TokenIntLiteral, ValueInt: 456}, {Type: TokenIntLiteral, ValueInt: 99}},
		},
		{
			input:          strings.NewReader(`ident`),
			expectedTokens: []Token{{Type: TokenIdent, ValueString: "ident"}},
		},
		{
			input:          strings.NewReader(`ident_with_underscores`),
			expectedTokens: []Token{{Type: TokenIdent, ValueString: "ident_with_underscores"}},
		},
		{
			input:          strings.NewReader(`true`),
			expectedTokens: []Token{{Type: TokenBool, ValueBool: true}},
		},
		{
			input:          strings.NewReader(`false`),
			expectedTokens: []Token{{Type: TokenBool, ValueBool: false}},
		},
		{
			input:          strings.NewReader(`ident1 ident2 true false`),
			expectedTokens: []Token{{Type: TokenIdent, ValueString: "ident1"}, {Type: TokenIdent, ValueString: "ident2"}, {Type: TokenBool, ValueBool: true}, {Type: TokenBool, ValueBool: false}},
		},
	}

	for i, c := range testCases {
		lexer := New(&characterReader{reader: c.input})

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

func Test_LexerHandler_Peek(t *testing.T) {
	testCases := []struct {
		input           io.Reader
		expectedToken   Token
		inputTokenCount int
	}{
		{
			input:           strings.NewReader(`"hello world"`),
			expectedToken:   Token{Type: TokenStringLiteral, ValueString: "hello world"},
			inputTokenCount: 1,
		},
		{
			input:           strings.NewReader(`"hello world""something""more text"`),
			expectedToken:   Token{Type: TokenStringLiteral, ValueString: "hello world"},
			inputTokenCount: 2,
		},
		{
			input:           strings.NewReader(`"hello world"!{}`),
			expectedToken:   Token{Type: TokenStringLiteral, ValueString: "hello world"},
			inputTokenCount: 4,
		},
	}

	for i, c := range testCases {
		lexer := New(&characterReader{reader: c.input})

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
