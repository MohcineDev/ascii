package handleArgs

import (
	"errors"
	"path"
	"regexp"

	"example.moh/handleFlag"
)

func CheckArgs(myArgs []string) (error, []string) {
	isOutput := false
	isColor := false
	color := ""
	outputFile := ""

	if len(myArgs) < 1 || len(myArgs) > 3 {
		return colorUsageMessage(), []string{}
	}
	///flag is output or color
	isOutput, outputFile, isColor, color = handleFlag.IsValidFlag(myArgs)

	///////////  OUTPUT ////////
	if len(myArgs) == 1 {
		if isOutput {
			return outputUsageMessage(), []string{}
		} else if isColor {
			///use color usage msg
			return colorUsageMessage(), []string{}
		} else if checkForDash(myArgs[0]) {
			if checkForFlagType(myArgs[0]) {
				return outputUsageMessage(), []string{}
			} else {
				return colorUsageMessage(), []string{}
			}
		} else {
			return nil, myArgs
		}
	} else if len(myArgs) == 2 {
		if isOutput {
			if len(outputFile) < 1 {
				return outputUsageMessage(), []string{}
			}
			myArgs = append(myArgs, "validFlag")
		} else if isColor {
			///color flag
			// if !handleFlag.HasEqualSign {
			// 	return colorUsageMessage(), []string{}
			// }
			if len(color) < 1 {
				///Error : color not found!!!
				return errors.New("CheckArgs Error : Color not found!!"), []string{}
			}

			myArgs = append(myArgs, "colorFlag")

		} else if checkForDash(myArgs[0]) {
			if checkForFlagType(myArgs[0]) {
				return outputUsageMessage(), []string{}
			} else {
				return colorUsageMessage(), []string{}
			}
		} else {
			///there is only one dash
			myArgs[1] = getBannerFileName(myArgs[1])
		}
	} else if len(myArgs) == 3 {
		if isOutput {
			myArgs = append(myArgs, "validFlag")
			if len(outputFile) < 1 {
				return outputUsageMessage(), []string{}
			}
			myArgs[2] = getBannerFileName(myArgs[2])
			///color flag
		} else if isColor {

			if len(color) < 1 {
				///Error : color not found!!!
				return errors.New("CheckArgs Error : Color not found!!"), []string{}
			}
			myArgs = append(myArgs, "colorFlag")

		}
	}

	return nil, myArgs
}

func colorUsageMessage() error {
	return errors.New("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
}

func outputUsageMessage() error {
	return errors.New("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
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
func checkForFlagType(flag string) bool {
	r, _ := regexp.Compile("output")

	return r.MatchString(flag)
}
