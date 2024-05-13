package handleFlag

import (
	"flag"
	"fmt"
)

var outputPtr = flag.String("output", "", "a string")

// // check if the flag is valid and return the file name
func IsValidFlag(myFlag string) (bool, string) {
	isValid := false

	fmt.Println("myFlag", myFlag)
	// fileExt := path.Ext(myFlag)

	outputFile := *outputPtr
	//handle out of range
	if len(myFlag) >= 8 && myFlag[:8] == "--output" {
		if len(myFlag) >= 9 {

			outputFile = myFlag[9:]
		}
		// if len(outputFile) >= 1 {
		isValid = true
		// }
	}
	return isValid, outputFile
}
