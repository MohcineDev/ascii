package handleArgs

import (
	"errors"
	"fmt"
	"path"
	"regexp"

	"example.moh/handleFlag"
)

var usageMsgs = map[string]error{
	"color":  errors.New("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\""),
	"output": errors.New("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"),
	"align":  errors.New("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard"),
}

func CheckArgs(myArgs []string) (error, []string) {
	if len(myArgs) < 1 || len(myArgs) > 3 {
		return usageMsgs["align"], []string{}
	}
	///flag is output or color
	isOutput, outputFile, isColor, color := handleFlag.IsValidFlag(myArgs)

	///////////  OUTPUT ////////
	if len(myArgs) == 1 {
		if isOutput {
			return usageMsgs["output"], []string{}
		} else if isColor {
			///use color usage msg
			return usageMsgs["color"], []string{}
		} else if handleFlag.IsAlign {
			////chck for align flag
			return usageMsgs["align"], []string{}
		} else if checkForDash(myArgs[0]) {
			////////displat usage msg based on the flag type
			if checkForFlagType(myArgs[0], "output") {
				return usageMsgs["output"], []string{}
			} else if checkForFlagType(myArgs[0], "color") {
				return usageMsgs["color"], []string{}
			} else {
				return usageMsgs["align"], []string{}
			}
		} else {
			return nil, myArgs
		}
	} else if len(myArgs) == 2 {
		if isOutput {
			if len(outputFile) < 1 {
				return usageMsgs["output"], []string{}
			}
		} else if isColor {
			///color flag
			if len(color) < 1 {
				///Error : color not found!!!
				return errors.New("CheckArgs Error : Color not found"), []string{}
			}
		} else if handleFlag.IsAlign {
			if len(handleFlag.Alignment) < 1 {
				fmt.Println("hello")
				return usageMsgs["align"], []string{}
			}
		} else if checkForDash(myArgs[0]) {
			if checkForFlagType(myArgs[0], "output") {
				return usageMsgs["output"], []string{}
			} else if checkForFlagType(myArgs[0], "color") {
				return usageMsgs["color"], []string{}
			} else {
				return usageMsgs["align"], []string{}
			}
		} else {
			myArgs[1] = getBannerFileName(myArgs[1])
		}
	} else if len(myArgs) == 3 {
		if isOutput {
			if len(outputFile) < 1 {
				return usageMsgs["output"], []string{}
			}
			myArgs[2] = getBannerFileName(myArgs[2])
			///color flag
		} else if isColor {
			if len(color) < 1 {
				///Error : color not found!!!
				return errors.New("CheckArgs Error : Color not found"), []string{}
			}
		} else if checkForDash(myArgs[0]) {
			if checkForFlagType(myArgs[0], "output") {
				return usageMsgs["output"], []string{}
			} else if checkForFlagType(myArgs[0], "color") {
				return usageMsgs["color"], []string{}
			} else {
				return usageMsgs["align"], []string{}
			}
		}
	}
	fmt.Println("myArgs : ", myArgs)
	return nil, myArgs
}

func getBannerFileName(Banner string) string {
	if path.Ext(Banner) != ".txt" {
		Banner += ".txt"
	}
	return Banner
}

// /check if there is a dash at the beginning
func checkForDash(flag string) bool {
	r, _ := regexp.Compile("^-+")

	return r.MatchString(flag)
}

// check if the flag is output or color
func checkForFlagType(flag string, flagName string) bool {
	r, _ := regexp.Compile(flagName)

	return r.MatchString(flag)
}
