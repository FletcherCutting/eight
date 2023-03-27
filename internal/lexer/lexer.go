package lexer

const (
	TokenIdentifier = iota
	TokenString
	TokenInt
	TokenBool

	TokenStringLiteral
	TokenIntLiteral
	TokenBoolLiteral

	TokenBang

	TokenEOF
)

type Lexer interface {
	Get() (Token, error)
	Peek() (Token, error)
}

type Token struct {
	Type        int
	ValueString string
	ValueInt    int
	ValueBool   bool
}

var tokenReadableNames = map[int]string{
	TokenIdentifier:    "TokenIdentifier",
	TokenString:        "TokenString",
	TokenInt:           "TokenInt",
	TokenBool:          "TokenBool",
	TokenStringLiteral: "TokenStringLiteral",
	TokenIntLiteral:    "TokenIntLiteral",
	TokenBoolLiteral:   "TokenBoolLiteral",
	TokenBang:          "TokenBang",
	TokenEOF:           "TokenEOF",
}
