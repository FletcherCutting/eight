package lexer

import (
	"fmt"
	"strconv"
	"unicode"
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

	eof, character, err := lh.reader.PeekNonWhitespaceCharacter()

	if eof {
		return Token{Type: TokenEOF}, nil
	}

	if err != nil {
		return Token{}, err
	}

	var returnToken Token

	// check is number
	if unicode.IsNumber(character) {
		returnToken, err = lh.readIntLiteral()
		goto ENDPEEK
	}

	// check is ident

	// check special characters
	switch character {
	case '"':
		lh.reader.Next()
		returnToken, err = lh.readStringLiteral()
	case '!':
		lh.reader.Next()
		returnToken = Token{Type: TokenBang}
	case '{':
		lh.reader.Next()
		returnToken = Token{Type: TokenOpenBrace}
	case '}':
		lh.reader.Next()
		returnToken = Token{Type: TokenCloseBrace}
	default:
		return Token{}, fmt.Errorf("unknown character: %v", string(character))
	}

ENDPEEK:
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
		eof, character, err := lh.reader.Read()

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

func (lh *LexerHandler) readIntLiteral() (Token, error) {
	intString := ""

	for {
		eof, character, err := lh.reader.Peek()

		if eof {
			break
		}

		if err != nil {
			return Token{}, err
		}

		if !unicode.IsNumber(character) {
			break
		}

		intString += string(character)
		lh.reader.Next()
	}

	returnInt, err := strconv.Atoi(intString)

	if err != nil {
		return Token{}, err
	}

	return Token{Type: TokenIntLiteral, ValueInt: returnInt}, nil
}
