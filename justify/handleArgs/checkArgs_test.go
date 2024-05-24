package handleArgs

import (
	"testing"
)

// /test with no arguments
func TestCheckArgsNoArg(t *testing.T) {
	runTest(t, []string{})
}

// /test with one arguments
func TestCheckArgsOneArg(t *testing.T) {
	runTest(t, []string{""})
}

// /test with two arguments
func TestCheckArgsTwoArgs(t *testing.T) {
	runTest(t, []string{"de", "ddf"})
}

// /test with three arguments
func TestCheckArgsThreeArgs(t *testing.T) {
	runTest(t, []string{"zz", "daz", "fsd"})
}

// /test with More thann three arguments
func TestCheckArgsMoreThanThreeArgs(t *testing.T) {
	runTest(t, []string{"zz", "daz", "fsd", "ds"})
}

func runTest(t *testing.T, args []string) {
	err, _ := CheckArgs(args)

	if err != nil {
		t.Fatalf(err.Error())
	}
}
