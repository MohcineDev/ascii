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

	outputFile := ""
	color := ""

	if len(myArgs) < 1 || len(myArgs) > 3 {
		return usageMessage(), []string{}
	}
	///flag is output or color
	isOutput, outputFile, isColor, color = handleFlag.IsValidFlag(myArgs)
	if isColor {
		fmt.Println(color)
	}
	if len(myArgs) == 1 {
		if isOutput {
			return usageMessage(), []string{}
		} else {
			return nil, myArgs
		}
	} else if len(myArgs) == 2 {
		if isOutput {
			if len(outputFile) < 1 {
				return usageMessage(), []string{}
			}
			myArgs = append(myArgs, "validFlag")
			fmt.Println("12")
		} else {
			myArgs[1] = getBannerFileName(myArgs[1])
		}
	} else if len(myArgs) == 3 {
		if isOutput {
			myArgs = append(myArgs, "validFlag")
			if len(outputFile) < 1 {
				return usageMessage(), []string{}
			}
			myArgs[2] = getBannerFileName(myArgs[2])
		}
	}

	return nil, myArgs
}

func usageMessage() error {
	return errors.New("usageMessage()\nUsage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
}

func getBannerFileName(Banner string) string {
	if path.Ext(Banner) != ".txt" {
		Banner += ".txt"
	}
	return Banner
}
