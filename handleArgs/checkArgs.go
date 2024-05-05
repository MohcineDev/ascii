package handleArgs

import (
	"errors"
	"os"
	"strings"
)

func CheckArgs() (error, []string) {
	myargs := os.Args[1:]

	if len(myargs) < 1 {
		return usageMessage(), []string{}
	} else if len(myargs) > 2 {
		return usageMessage(), []string{}
	}
	if len(myargs) == 2 {

		Banner := myargs[1]

		if strings.Contains(Banner, ".") {
			if Banner == "standard.txt" || Banner == "standard" {
				return nil, []string{myargs[0], "standard"}
			} else if Banner == "shadow.txt" || Banner == "shadow" {
				return nil, []string{myargs[0], "shadow"}
			} else if Banner == "thinkertoy.txt" || Banner == "thinkertoy" {
				return nil, []string{myargs[0], "thinkertoy"}
			} else if Banner == "mine.txt" || Banner == "mine" {
				return nil, []string{myargs[0], "mine"}
			}
		}
	}
	return nil, myargs
}

func usageMessage() error {
	return errors.New(`Usage: go run . [STRING] [BANNER]

EX: go run . something standard`)
}
