package lexer

import (
	"fmt"
)

type LexerHandler struct {
	reader   *characterReader
	tokens   []Token
	position int
}

func New(reader *characterReader) *LexerHandler {
	return &LexerHandler{
		reader:   reader,
		position: -1,
	}
}

func (lh *LexerHandler) Read() (Token, error) {
	if lh.positionBehindTokensLength() {
		token := lh.tokens[lh.position+1]
		lh.position++
		return token, nil
	}

	token, err := lh.Peek()

	if err != nil {
		return Token{}, fmt.Errorf("failed to read token: %v", err)
	}

	lh.position++
	return token, nil
}

func (lh *LexerHandler) Peek() (Token, error) {
	if lh.positionBehindTokensLength() {
		return lh.tokens[lh.position+1], nil
	}

	eof, character, err := lh.reader.getNextNonWhitespaceCharacter()

	if eof {
		return Token{Type: TokenEOF}, nil
	}

	if err != nil {
		return Token{}, err
	}

	var returnToken Token

	switch character {
	case '"':
		returnToken, err = lh.readStringLiteral()
	case '!':
		returnToken = Token{Type: TokenBang}
	case '{':
		returnToken = Token{Type: TokenOpenBrace}
	case '}':
		returnToken = Token{Type: TokenCloseBrace}
	default:
		return Token{}, fmt.Errorf("unknown character: %v", character)
	}

	if err != nil {
		return Token{}, err
	}

	lh.tokens = append(lh.tokens, returnToken)

	return returnToken, nil
}

func (lh *LexerHandler) positionBehindTokensLength() bool {
	tokensLength := len(lh.tokens)
	return tokensLength > 0 && tokensLength-1 > lh.position
}

func (lh *LexerHandler) readStringLiteral() (Token, error) {
	returnString := ""

	for {
		eof, character, err := lh.reader.getNextCharacter()

		if eof {
			return Token{}, fmt.Errorf("unexpected EOF")
		}

		if err != nil {
			return Token{}, err
		}

		if character == '"' {
			break
		}

		returnString += string(character)
	}

	return Token{Type: TokenStringLiteral, ValueString: returnString}, nil
}
