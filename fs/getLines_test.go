package main

import (
	"os"
	"testing"
)

////////RUN TEST :  go test *.go -v

////test file

// /no file provided should word // it uses standard file by default
func TestGetLinesWithoutFile(t *testing.T) {
	os.Args = []string{"go run .", "fsd"}

	line, _ := getLines()

	// line not found // empty
	if line == nil {
		t.Fatalf(`something here error`)
	}
}

// //wrong file name // shold return not found

// //wrong file name // shold return not found
func TestGetLinesWrongFile(t *testing.T) {
	os.Args = []string{"go run .", "fsd", "shadowwww"}

	line, _ := getLines()

	// line not found // empty
	if line == nil {
		t.Fatalf(`something here error`)
	}
}

func TestGetLinesMoreArgs(t *testing.T) {
	os.Args = []string{"go run .", "fsd", "shadow", "e"}

	line, _ := getLines()

	// line not found // empty
	if line == nil {
		t.Fatalf(`something here error`)
	}
}
