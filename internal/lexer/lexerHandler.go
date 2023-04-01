package lexer

import (
	"fmt"
	"io"
	"unicode"
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

	eof, character, err := lh.getNextNonWhitespaceCharacter()

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

func (lh *LexerHandler) getNextCharacter() (bool, rune, error) {
	characterBuffer := make([]byte, 1)
	bytesRead, err := lh.data.Read(characterBuffer)

	if bytesRead == 0 && err == io.EOF {
		return true, 0, nil
	}

	if err != nil {
		return false, 0, err
	}

	if bytesRead == 0 {
		return false, 0, fmt.Errorf("failed to read bytes")
	}

	return false, rune(characterBuffer[0]), nil
}

func (lh *LexerHandler) getNextNonWhitespaceCharacter() (bool, rune, error) {
	var returnCharacter rune

	for {
		characterBuffer := make([]byte, 1)

		bytesRead, err := lh.data.Read(characterBuffer)

		if bytesRead == 0 && err == io.EOF {
			return true, 0, nil
		}

		if err != nil {
			return false, 0, fmt.Errorf("errored when reading bytes: %v", err)
		}

		if bytesRead == 0 {
			return false, 0, fmt.Errorf("failed to read bytes")
		}

		returnCharacter = rune(characterBuffer[0])

		if !unicode.IsSpace(returnCharacter) {
			break
		}
	}

	return false, returnCharacter, nil
}

func (lh *LexerHandler) positionBehindTokensLength() bool {
	tokensLength := len(lh.tokens)
	return tokensLength > 0 && tokensLength-1 > lh.position
}

func (lh *LexerHandler) readStringLiteral() (Token, error) {
	returnString := ""

	for {
		eof, character, err := lh.getNextCharacter()

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
