package handleArgs

import (
	"testing"
)

// /test with any args
func TestCheckArgsNoArg(t *testing.T) {
	myArgs := []string{}
	err, _ := CheckArgs(myArgs)

	if err != nil {
		t.Fatalf(`CheckArgs() = Usage: go run . [STRING] [BANNER] error`)
	}
}

// /test with one arg
func TestCheckArgsOneArg(t *testing.T) {
	myArgs := []string{""}
	err, _ := CheckArgs(myArgs)

	if err != nil {
		t.Fatalf(`CheckArgs() = Usage: go run . [STRING] [BANNER] error`)
	}
}

// /test with two args
func TestCheckArgsTwoArgs(t *testing.T) {
	myArgs := []string{"de", "ddf"}
	err, _ := CheckArgs(myArgs)

	if err != nil {
		t.Fatalf(`CheckArgs() = Usage: go run . [STRING] [BANNER] error`)
	}
}

// /test with out more than two args
func TestCheckArgsMoreThanTwoArgs(t *testing.T) {
	myArgs := []string{"zz", "daz", "fsd"}
	err, _ := CheckArgs(myArgs)

	if err != nil {
		t.Fatalf(`CheckArgs() = Usage: go run . [STRING] [BANNER] error`)
	}
}
