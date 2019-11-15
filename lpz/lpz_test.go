package lpz

import (
	"bytes"
	"io"
	"testing"
)

func TestLexicon_MaxIndexSize(t *testing.T) {
	lex := Lexicon{
		size: 2,
		root: nil,
	}
	if lex.MaxIndexSize() != 1 {
		t.Error() // to indicate test failed
	}

	lex.size = 3
	if lex.MaxIndexSize() != 2 {
		t.Error() // to indicate test failed
	}

}

func bufIsEquals(arr bytes.Buffer, str string) bool {
	return arr.String() == str
}

func TestWriteInFile(t *testing.T) {
	var arr bytes.Buffer
	WriteInFile(io.Writer(&arr), 5, 4)
	if !bufIsEquals(arr, "0101") {
		t.Error()
	}
	arr.Reset()
	WriteInFile(io.Writer(&arr), 5, 3)
	if !bufIsEquals(arr, "101") {
		t.Error()
	}
}
