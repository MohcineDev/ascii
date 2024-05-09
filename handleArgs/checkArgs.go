package handleArgs

import (
	"errors"
	"handleFlag"
)

func CheckArgs(myArgs []string) (error, []string) {
	/*
		if <1 +>>>> err msg
		if == 2 add func if valid flag
			//y : save the result to the file using standard file format
			//no : show err flag not right
			/// check the second one if it is a valid file
		if == 3
	*/

	if len(myArgs) < 1 {
		return usageMessage(), []string{}
	} else if len(myArgs) == 2 {
		isValid, _ := handleFlag.IsValidFlag()

		if isValid {
			myArgs = append(myArgs, "validFlag")
		}
	} else if len(myArgs) == 3 {
		isValid, _ := handleFlag.IsValidFlag()
		if isValid {
			myArgs = append(myArgs, "validFlag") 

		}
	}

	//// fs project

	// if the file is provided
	/*	if len(myArgs) == 2 {

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
	}*/
	return nil, myArgs

	/*
		//////////////  USED IN THE FS PROJECT
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
	*/
}

func usageMessage() error {
	return errors.New(`Usage: go run . [STRING] [BANNER]

EX: go run . something standard`)
}
