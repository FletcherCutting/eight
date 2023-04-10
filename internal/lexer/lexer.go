package lexer

const (
	TokenUnknown tokenType = iota
	TokenIdent
	TokenString
	TokenInt
	TokenBool

	TokenStringLiteral
	TokenIntLiteral
	TokenBoolLiteral

	TokenBang
	TokenOpenBrace
	TokenCloseBrace

	TokenEOF
)

type tokenType int

func (tt tokenType) String() string {
	return tokenReadableNames[tt]
}

type Lexer interface {
	Get() (Token, error)
	Peek() (Token, error)
}

type Token struct {
	Type        tokenType
	ValueString string
	ValueInt    int
	ValueBool   bool
}

var tokenReadableNames = map[tokenType]string{
	TokenUnknown:       "TokenUnknown",
	TokenIdent:         "TokenIdent",
	TokenString:        "TokenString",
	TokenInt:           "TokenInt",
	TokenBool:          "TokenBool",
	TokenStringLiteral: "TokenStringLiteral",
	TokenIntLiteral:    "TokenIntLiteral",
	TokenBoolLiteral:   "TokenBoolLiteral",
	TokenBang:          "TokenBang",
	TokenOpenBrace:     "TokenOpenBrace",
	TokenCloseBrace:    "TokenCloseBrace",
	TokenEOF:           "TokenEOF",
}

var knownIdents = map[string]Token{
	"true":  {Type: TokenBool, ValueBool: true},
	"false": {Type: TokenBool, ValueBool: false},
}
