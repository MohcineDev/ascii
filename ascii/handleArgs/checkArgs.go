package handleArgs

import (
	"errors"
//	"fmt"
	"os"
)

func CheckArgs() (error, string) {
	myargs := os.Args[1:]
	if len(myargs) < 1 {
		return errors.New("\n--Error please enter an input!!\n"), ""
	} else if len(myargs) > 1 {
		return errors.New("\n--Error: please enter one input only!!\n"), ""
	}

	return nil, myargs[0]
}
