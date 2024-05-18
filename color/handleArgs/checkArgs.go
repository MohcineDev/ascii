package handleArgs

import (
	"errors"
	"fmt"
	"path"

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
	fmt.Println("FLAGS : ", isOutput, isColor)
	///////////  OUTPUT ////////
	if len(myArgs) == 1 {
		if isOutput {
			return outputUsageMessage(), []string{}
		} else if isColor {
			///use color usage msg
			return colorUsageMessage(), []string{}
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

			if len(color) < 1 {
				///Error : color not found!!!
				return errors.New("Error : Color not found!!"), []string{}
			}
			myArgs = append(myArgs, "colorFlag")

		} else {
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
				return errors.New("Error : Color not found!!"), []string{}
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
