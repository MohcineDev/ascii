package handleArgs

import (
	"errors"
)

func CheckArgs(myArgs []string) (error, []string) {
	if len(myArgs) < 1 || len(myArgs) > 2 {
		return usageMessage(), []string{}
	}
	//// fs project

	// if the file is provided
	if len(myArgs) == 2 {

		Banner := myArgs[1]
 
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
	return nil, myArgs
}

func usageMessage() error {
	return errors.New(`Usage: go run . [STRING] [BANNER]

EX: go run . something standard`)
}
