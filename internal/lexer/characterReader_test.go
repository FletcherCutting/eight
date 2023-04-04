package lexer

import (
	"io"
	"strings"
	"testing"
)

func Test_CharacterReader_GetNextCharacter(t *testing.T) {
	testCases := []struct {
		reader             io.Reader
		loopTimes          int
		expectedEOFs       []bool
		expectedCharacters []rune
		expectedErrors     []error
	}{
		{
			reader:             strings.NewReader("abcd"),
			loopTimes:          4,
			expectedEOFs:       []bool{false, false, false, false},
			expectedCharacters: []rune{'a', 'b', 'c', 'd'},
			expectedErrors:     []error{nil, nil, nil, nil},
		},
		{
			reader:             strings.NewReader("123!@#"),
			loopTimes:          6,
			expectedEOFs:       []bool{false, false, false, false, false, false},
			expectedCharacters: []rune{'1', '2', '3', '!', '@', '#'},
			expectedErrors:     []error{nil, nil, nil, nil, nil, nil},
		},
		{
			reader:             strings.NewReader("a c\nd\te"),
			loopTimes:          7,
			expectedEOFs:       []bool{false, false, false, false, false, false, false},
			expectedCharacters: []rune{'a', ' ', 'c', '\n', 'd', '\t', 'e'},
			expectedErrors:     []error{nil, nil, nil, nil, nil, nil, nil},
		},
		{
			reader:             strings.NewReader("abcd"),
			loopTimes:          5,
			expectedEOFs:       []bool{false, false, false, false, true},
			expectedCharacters: []rune{'a', 'b', 'c', 'd'},
			expectedErrors:     []error{nil, nil, nil, nil},
		},
	}

	for i, tc := range testCases {
		cr := &characterReader{reader: tc.reader}

		for j := 0; j < tc.loopTimes; j++ {
			eof, character, err := cr.getNextCharacter()

			if eof != tc.expectedEOFs[j] {
				t.Errorf("case %d-%d: Unexpected EOF\nExpected: %v\nActual: %v", i, j, tc.expectedEOFs[j], eof)
			}

			if eof && tc.expectedEOFs[j] {
				continue
			}

			if err != tc.expectedErrors[j] {
				t.Errorf("case %d-%d: Unexpected error\nExpected: %v\nActual: %v", i, j, tc.expectedErrors[j], err)
			}

			if character != tc.expectedCharacters[j] {
				t.Errorf("case %d-%d: Unexpected character\nExpected: %s\nActual: %s", i, j, string(tc.expectedCharacters[j]), string(character))
			}
		}
	}
}

func Test_CharacterReader_GetNextNonWhitespaceCharacter(t *testing.T) {
	testCases := []struct {
		reader             io.Reader
		loopTimes          int
		expectedEOFs       []bool
		expectedCharacters []rune
		expectedErrors     []error
	}{
		{
			reader:             strings.NewReader("123!@#"),
			loopTimes:          6,
			expectedEOFs:       []bool{false, false, false, false, false, false},
			expectedCharacters: []rune{'1', '2', '3', '!', '@', '#'},
			expectedErrors:     []error{nil, nil, nil, nil, nil, nil},
		},
	}

	for i, tc := range testCases {
		cr := &characterReader{reader: tc.reader}

		for j := 0; j < tc.loopTimes; j++ {
			eof, character, err := cr.getNextNonWhitespaceCharacter()

			if eof != tc.expectedEOFs[j] {
				t.Errorf("case %d-%d: Unexpected EOF\nExpected: %v\nActual: %v", i, j, tc.expectedEOFs[j], eof)
			}

			if eof && tc.expectedEOFs[j] {
				continue
			}

			if err != tc.expectedErrors[j] {
				t.Errorf("case %d-%d: Unexpected error\nExpected: %v\nActual: %v", i, j, tc.expectedErrors[j], err)
			}

			if character != tc.expectedCharacters[j] {
				t.Errorf("case %d-%d: Unexpected character\nExpected: %s\nActual: %s", i, j, string(tc.expectedCharacters[j]), string(character))
			}
		}
	}
}

func Test_CharacterReader_Peek(t *testing.T) {
	testString := "abcd"
	cr := &characterReader{reader: strings.NewReader(testString)}
	expectedCharacter := 'a'

	for i := range testString {
		eof, character, err := cr.Peek()

		if eof {
			t.Errorf("%d: Unexpected eof", i)
		}

		if err != nil {
			t.Errorf("%d: Unexpected error: %v", i, err)
		}

		if character != expectedCharacter {
			t.Errorf("%d: Unexpected character\nExpected: %s\nActual: %s", i, string(expectedCharacter), string(character))
		}
	}
}

func Test_CharacterReader_PeekNonWhitespaceCharacter(t *testing.T) {
	testString := "  \n\tabcd"
	cr := &characterReader{reader: strings.NewReader(testString)}
	expectedCharacter := 'a'

	for i := range testString {
		eof, character, err := cr.PeekNonWhitespaceCharacter()

		if eof {
			t.Errorf("%d: Unexpected eof", i)
		}

		if err != nil {
			t.Errorf("%d: Unexpected error: %v", i, err)
		}

		if character != expectedCharacter {
			t.Errorf("%d: Unexpected character\nExpected: %s\nActual: %s", i, string(expectedCharacter), string(character))
		}
	}
}
