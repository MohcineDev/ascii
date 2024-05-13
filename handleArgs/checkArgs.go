package handleArgs

import (
	"errors"
	"fmt"
	"path"

	"example.moh/handleFlag"
)

func CheckArgs(myArgs []string) (error, []string) {
	/*
		if < 0 || > 2 +>>>> err msg
		if == 1 add func if the flag is valid  //--- go run . --output=result.txt "HELLO"
			//y : save the result to the result.txt file using standard file format
			//no : display the text in the terminal using the standard format
		if == 2
			//check if it's a valid flag
				/// y : check if the second arg is a valid file format name
				 	/// y : save the result to the result.txt file using specified file format
					/// n : display err file not found
				///n : display err file not found
	*/

	if len(myArgs) < 1 || len(myArgs) > 3 {
		return usageMessage(), []string{}
	}
	isValid, _ := handleFlag.IsValidFlag(myArgs[0])
	fmt.Println(isValid, myArgs[0])
	if len(myArgs) == 1 {
		/// if the flag is valid save the text to the result file using standard file format
		// if string(myArgs[0][0]) == "-" && string(myArgs[0][1]) != "-" {
		// 	return usageMessage(), []string{}
		// }
		if isValid {
			fmt.Println("123")
			return usageMessage(), []string{}
		} else {
			return nil, myArgs
		}
	} else if len(myArgs) == 2 {
		fmt.Println(myArgs)
		if string(myArgs[0][0]) == "-" && string(myArgs[0][1]) != "-" {
			return usageMessage(), []string{}
		}
		if isValid {
			myArgs = append(myArgs, "validFlag")
		} else {
			myArgs[1] = getBannerFileName(myArgs[1])
		}
	} else if len(myArgs) == 3 {
		if isValid {
			myArgs = append(myArgs, "validFlag")

			myArgs[2] = getBannerFileName(myArgs[2])
		} else {
		}
	}

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
	if path.Ext(Banner) != ".txt" {
		Banner += ".txt"
	}
	return Banner
}

func startWithDash() {
}
