package handleArgs

import (
	"errors"
	"fmt"
	"path"
	"strings"

	"example.moh/handleFlag"
)

func CheckArgs(myArgs []string) (error, []string) {
	/*
		if <1 +>>>> err msg
		if == 2 add func if the flag is valid  //--- go run . --output=result.txt "HELLO"
			//y : save the result to the result.txt file using standard file format
			//no : check if the second arg is a valid file format name
				/// y : display the text in the terminal using the specified format
				/// n : display err file not found
		if == 3
			/// check the second one if it is a valid file
	*/

	if len(myArgs) < 1 || len(myArgs) > 3 {
		return usageMessage(), []string{}
	} else if len(myArgs) == 2 {
		isValid, _ := handleFlag.IsValidFlag()

		/// if the flag is valid save the text to the result file using standard file format
		if isValid {
			myArgs = append(myArgs, "validFlag")
		} else {
			myArgs[1] = getBannerFileName(myArgs[1])
		}
	} else if len(myArgs) == 3 {

		isValid, _ := handleFlag.IsValidFlag()

		if isValid {
			if strings.Contains(myArgs[2], ".") {
				myArgs[2] = getBannerFileName(myArgs[2])
			}
			myArgs = append(myArgs, "validFlag")
		}

	}

	fmt.Println("end : ", myArgs)
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
	// error strings should not be capitalized
	return errors.New("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
}

func getBannerFileName(Banner string) string {
	fileName := ""
	if path.Ext(Banner) != ".txt" {
		fileName = Banner + ".txt"
	}
	fmt.Println("hi from getBannerFileName!!!!!!!!!")

	return fileName
}
