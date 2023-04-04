package lexer

import (
	"fmt"
	"io"
	"unicode"
)

type characterReader struct {
	reader            io.Reader
	bufferedCharacter *rune
}

func (cr *characterReader) Peek() (bool, rune, error) {
	if cr.bufferedCharacter != nil {
		return false, *cr.bufferedCharacter, nil
	}

	eof, character, err := cr.getNextCharacter()

	if eof {
		return true, 0, nil
	}

	if err != nil {
		return false, 0, err
	}

	cr.bufferedCharacter = &character

	return false, character, nil
}

func (cr *characterReader) PeekNonWhitespaceCharacter() (bool, rune, error) {
	if cr.bufferedCharacter != nil && !unicode.IsSpace(*cr.bufferedCharacter) {
		return false, *cr.bufferedCharacter, nil
	}

	eof, character, err := cr.getNextNonWhitespaceCharacter()

	if eof {
		return true, 0, nil
	}

	if err != nil {
		return false, 0, err
	}

	cr.bufferedCharacter = &character

	return false, character, nil
}

func (cr *characterReader) Read() (bool, rune, error) {
	if cr.bufferedCharacter != nil {
		character := *cr.bufferedCharacter
		cr.bufferedCharacter = nil
		return false, character, nil
	}

	eof, character, err := cr.getNextCharacter()

	if eof {
		return true, 0, nil
	}

	if err != nil {
		return false, 0, err
	}

	return false, character, nil
}

func (cr *characterReader) ReadNonWhitespaceCharacter() (bool, rune, error) {
	if cr.bufferedCharacter != nil && !unicode.IsSpace(*cr.bufferedCharacter) {
		character := *cr.bufferedCharacter
		cr.bufferedCharacter = nil
		return false, character, nil
	}

	eof, character, err := cr.getNextNonWhitespaceCharacter()

	if eof {
		return true, 0, nil
	}

	if err != nil {
		return false, 0, err
	}

	return false, character, nil
}

func (cr *characterReader) getNextCharacter() (bool, rune, error) {
	characterBuffer := make([]byte, 1)
	bytesRead, err := cr.reader.Read(characterBuffer)

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

func (cr *characterReader) getNextNonWhitespaceCharacter() (bool, rune, error) {
	for {
		eof, character, err := cr.getNextCharacter()

		if eof {
			return true, 0, nil
		}

		if err != nil {
			return false, 0, err
		}

		if !unicode.IsSpace(character) {
			return false, character, nil
		}
	}
}
