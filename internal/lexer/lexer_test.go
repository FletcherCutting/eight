package lexer

import (
	"strings"
	"testing"
)

func TestPeek(t *testing.T) {
	input := strings.NewReader("\"hello world\"\"something\"")
	l := New(input)
	token, err := l.Peek()

	expectedType := TokenStringLiteral
	expectedValueString := "hello world"

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if token.Type != expectedType {
		t.Fatalf("Token.Type not as expected\nExpected: %s\nActual: %s", expectedType, token.Type)
	}

	if token.ValueString != expectedValueString {
		t.Fatalf("Token.ValueString not as expected\nExpected: %q\nActual: %q", expectedValueString, token.ValueString)
	}

	token, err = l.Peek()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if token.Type != expectedType {
		t.Fatalf("Token.Type not as expected\nExpected: %s\nActual: %s", expectedType, token.Type)
	}

	if token.ValueString != expectedValueString {
		t.Fatalf("Token.ValueString not as expected\nExpected: %q\nActual: %q", expectedValueString, token.ValueString)
	}
}
