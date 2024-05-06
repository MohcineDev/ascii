package handleArgs

import (
	"fmt"
	"os"
	"testing"
)

func TestCheckArgs(t *testing.T) {
	os.Args = append(os.Args, "sd")
	os.Args = append(os.Args, "standard")
	fmt.Println(os.Args)
	msg, err := CheckArgs()
	if err != nil {
		t.Fatalf(`CheckArgs() = %q, %v, want "", error`, msg, err)
	}
}
