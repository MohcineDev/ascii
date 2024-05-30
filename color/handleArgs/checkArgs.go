package handleArgs

import (
	"errors"
	"regexp"

	"example.moh/handleFlag"
)

var usageMsgs = map[string]error{
	"color":  errors.New("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\""),
	"output": errors.New("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"),
}

func CheckArgs(myArgs []string) (error, []string) {
	if len(myArgs) < 1 || len(myArgs) > 3 {
		return usageMsgs["color"], []string{}
	}
	handleFlag.IsValidFlag(myArgs)

	///////////  OUTPUT ////////
	if len(myArgs) == 1 {
		if handleFlag.IsOutput {
			return usageMsgs["output"], []string{}
		} else if handleFlag.IsColor {
			///use color usage msg
			return usageMsgs["color"], []string{}
		} else if checkForDash(myArgs[0]) {
			if checkForFlagType(myArgs[0]) {
				return usageMsgs["output"], []string{}
			} else {
				return usageMsgs["color"], []string{}
			}
		} else {
			return nil, myArgs
		}
	} else if len(myArgs) == 2 {
		if handleFlag.IsOutput {
			if len(handleFlag.OutputFile) < 1 {
				return usageMsgs["output"], []string{}
			}
		} else if handleFlag.IsColor {
			if len(handleFlag.Color) < 1 {
				///Error : color not found!!!
				return usageMsgs["color"], []string{}

			}
		} else if checkForDash(myArgs[0]) {
			if checkForFlagType(myArgs[0]) {
				return usageMsgs["output"], []string{}
			} else {
				return usageMsgs["color"], []string{}
			}
		} else {
			myArgs[1] = getBannerFileName(myArgs[1])
		}
	} else if len(myArgs) == 3 {
		if handleFlag.IsOutput {
			if len(handleFlag.OutputFile) < 1 {
				return usageMsgs["output"], []string{}
			}
			myArgs[2] = getBannerFileName(myArgs[2])
			///color flag
		} else if handleFlag.IsColor {
			if len(handleFlag.Color) < 1 {
				///Error : color not found!!!
				return usageMsgs["color"], []string{}

			}
		} else if checkForDash(myArgs[0]) {
			if checkForFlagType(myArgs[0]) {
				return usageMsgs["output"], []string{}
			} else {
				return usageMsgs["color"], []string{}
			}
		} else {
			return usageMsgs["color"], []string{}

		}
	}

	return nil, myArgs
}

func getBannerFileName(Banner string) string {
	return Banner + ".txt"
}

// /check if there is a dash at the beginning
func checkForDash(flag string) bool {
	r, _ := regexp.Compile("^--output|--color+")

	return r.MatchString(flag)
}

// check if the flag is output or color
func checkForFlagType(flag string) bool {
	r, _ := regexp.Compile("output")

	return r.MatchString(flag)
}
