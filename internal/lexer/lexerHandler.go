package lexer

import (
	"fmt"
	"io"
)

type LexerHandler struct {
	data     io.Reader
	tokens   []Token
	position int
}

func New(data io.Reader) *LexerHandler {
	return &LexerHandler{
		data:     data,
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

	characterBuffer := make([]byte, 1)

	bytesRead, err := lh.data.Read(characterBuffer)

	if bytesRead == 0 && err == io.EOF {
		return Token{Type: TokenEOF}, nil
	}

	if err != nil {
		return Token{}, fmt.Errorf("errored when reading bytes: %v", err)
	}

	if bytesRead == 0 {
		return Token{}, fmt.Errorf("failed to read bytes")
	}

	character := rune(characterBuffer[0])
	var returnToken Token

	switch character {
	case '"':
		stringLiteral, err := lh.readStringLiteral()

		if err != nil {
			return Token{}, fmt.Errorf("failed when reading string literal: %v", err)
		}

		returnToken = Token{Type: TokenStringLiteral, ValueString: stringLiteral}
	default:
		return Token{}, fmt.Errorf("unknown character: %v", character)
	}

	lh.tokens = append(lh.tokens, returnToken)

	return returnToken, nil
}

func (lh *LexerHandler) positionBehindTokensLength() bool {
	tokensLength := len(lh.tokens)
	return tokensLength > 0 && tokensLength-1 > lh.position
}

func (lh *LexerHandler) readStringLiteral() (string, error) {
	returnString := ""
	characterBuffer := make([]byte, 1)

	for {
		bytesRead, err := lh.data.Read(characterBuffer)

		if err != nil {
			return "", fmt.Errorf("failed when reading string literal: %v", err)
		}

		if bytesRead == 0 {
			return "", fmt.Errorf("no bytes read")
		}

		character := rune(characterBuffer[0])

		if character == '"' {
			break
		}

		returnString += string(character)
	}

	return returnString, nil
}
