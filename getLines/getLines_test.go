package getLines

import (
	"os"
	"testing"
)

////////RUN TEST :  go test *.go -v /// bcs main package

////test file

// /no file provided should word // it uses standard file by default
func TestGetLinesWithoutFile(t *testing.T) {
	os.Args = []string{"go run .", "AA"}

	_, _, err := GetLines()
	// line not found // empty
	if err != nil {
		t.Fatalf(`error msg`)
	}
}



// //wrong file name // should return not found
func TestGetLinesWrongFile(t *testing.T) {
	os.Args = []string{"go run .", "FF", "shadowwww"}

	_, _, err := GetLines()
	// line not found // empty
	if err != nil {
		t.Fatalf(`error msg`)
	}
}
 

/// test with More Args
func TestGetLinesMoreArgs(t *testing.T) {
	os.Args = []string{"go run .", "DD", "shadow", "arg3"}

	_, _, err := GetLines()
	// line not found // empty
	if err != nil {
		t.Fatalf(`error msg`)
	}
}