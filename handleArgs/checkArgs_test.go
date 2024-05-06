package handleArgs

import (
	"testing"
)

///test with any args
func TestCheckArgsNoArg(t *testing.T) {
	myArgs := []string{}
	err, msg := CheckArgs(myArgs)

	if err != nil {
		t.Fatalf(`CheckArgs() = %q, %v, want "", error`, msg, err)
	}
}
///test with one arg
func TestCheckArgsOneArg(t *testing.T) {
	myArgs := []string{""}
	err, msg := CheckArgs(myArgs)

	if err != nil {
		t.Fatalf(`CheckArgs() = %q, %v, want "", error`, msg, err)
	}
}
///test with two args
func TestCheckArgsTwoArgs(t *testing.T) {
	myArgs := []string{}
	err, msg := CheckArgs(myArgs)

	if err != nil {
		t.Fatalf(`CheckArgs() = %q, %v, want "", error`, msg, err)
	}
}
///test with out more than two args
func TestCheckArgsMoreThanTwoArgs(t *testing.T) {
	myArgs := []string{}
	err, msg := CheckArgs(myArgs)

	if err != nil {
		t.Fatalf(`CheckArgs() = %q, %v, want "", error`, msg, err)
	}
}