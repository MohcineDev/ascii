package handleArgs

import (
	"errors"
	"regexp"

	"example.moh/handleFlag"
)

var usageMsgs = map[string]error{
	"color":  errors.New("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\""),
	"output": errors.New("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"),
	"align":  errors.New("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard"),
}

func CheckArgs(myArgs []string) ([]string, error) {
	if len(myArgs) < 1 || len(myArgs) > 4 {
		return []string{}, usageMsgs["align"]
	}
	///check the used flag
	handleFlag.IsValidFlag(myArgs)

	switch len(myArgs) {
	case 1:
		if handleFlag.IsOutput {
			return []string{}, usageMsgs["output"]
		} else if handleFlag.IsColor {
			///use color usage msg
			return []string{}, usageMsgs["color"]
		} else if handleFlag.IsAlign {
			////chck for align flag
			return []string{}, usageMsgs["align"]
		} else if checkForDash(myArgs[0]) {
			return []string{}, displayMsg(myArgs[0])
		} else {
			return myArgs, nil
		}
	case 2:

		if handleFlag.IsOutput {
			if len(handleFlag.OutputFile) < 1 {
				return []string{}, usageMsgs["output"]
			}
		} else if handleFlag.IsColor {
			///color flag
			if len(handleFlag.Color) < 1 {
				///Error : color not found!!!
				return []string{}, errors.New("CheckArgs Error : Color not found")
			}
		} else if handleFlag.IsAlign {
			if len(handleFlag.Alignment) < 1 || !checkAlignment(handleFlag.Alignment) {

				return []string{}, usageMsgs["align"]
			}
		} else if checkForDash(myArgs[0]) {
			return []string{}, displayMsg(myArgs[0])

		} else {
			myArgs[1] = getBannerFileName(myArgs[1])
		}
	case 3:
		if handleFlag.IsOutput {
			if len(handleFlag.OutputFile) < 1 {
				return []string{}, usageMsgs["output"]
			}
			myArgs[2] = getBannerFileName(myArgs[2])
			///color flag
		} else if handleFlag.IsColor {
			if len(handleFlag.Color) < 1 {
				///Error : color not found!!!
				return []string{}, errors.New("CheckArgs Error : Color not found")
			}
		} else if handleFlag.IsAlign {
			if len(handleFlag.Alignment) < 1 {
				///Error : color not found!!!
				return []string{}, errors.New("alignment  error ")
			}
			myArgs[2] = getBannerFileName(myArgs[2])
		} else if checkForDash(myArgs[0]) {
			return []string{}, displayMsg(myArgs[0])
		}
	case 4:
		if handleFlag.IsColor {
			if len(handleFlag.Color) < 1 {
				///Error : color not found!!!
				return []string{}, errors.New("CheckArgs Error : Color not found")
			}
			myArgs[3] = getBannerFileName(myArgs[3])

		}
	}

	return myArgs, nil
}

func getBannerFileName(Banner string) string {
	return Banner + ".txt"

}

// /check if there is a dash at the beginning of the flag
func checkForDash(flag string) bool {
	r, _ := regexp.Compile("^-+")
	return r.MatchString(flag)
}

// check if the flag is output or color
func checkForFlagType(flag string, flagName string) bool {
	r, _ := regexp.Compile(flagName)
	return r.MatchString(flag)
}

// /check the align flag value
func checkAlignment(value string) bool {
	if value == "center" || value == "right" || value == "justify" {
		return true
	}
	return false
}

// //////display usage msg based on the flag type
func displayMsg(flagArg string) error {

	if checkForFlagType(flagArg, "output") {
		return usageMsgs["output"]
	} else if checkForFlagType(flagArg, "color") {
		return usageMsgs["color"]
	} else {
		return usageMsgs["align"]
	}
}
