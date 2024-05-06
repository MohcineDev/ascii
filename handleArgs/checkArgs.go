package handleArgs

import (
	"errors"
	"strings"
)

func CheckArgs(myArgs []string) (error, []string) {
	if len(myArgs) < 1 {
		return usageMessage(), []string{}
	} else if len(myArgs) > 2 {
		return usageMessage(), []string{}
	}

	//// fs project

	if len(myArgs) == 2 {

		Banner := myArgs[1]

		if strings.Contains(Banner, ".") {
			if Banner == "standard.txt" || Banner == "standard" {
				return nil, []string{myArgs[0], "standard"}
			} else if Banner == "shadow.txt" || Banner == "shadow" {
				return nil, []string{myArgs[0], "shadow"}
			} else if Banner == "thinkertoy.txt" || Banner == "thinkertoy" {
				return nil, []string{myArgs[0], "thinkertoy"}
			} else if Banner == "mine.txt" || Banner == "mine" {
				return nil, []string{myArgs[0], "mine"}
			}
		}
	}
	return nil, myArgs
}

func usageMessage() error {
	return errors.New(`Usage: go run . [STRING] [BANNER]

EX: go run . something standard`)
}
