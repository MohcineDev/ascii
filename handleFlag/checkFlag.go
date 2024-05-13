package handleFlag

import (
	"flag"
	"fmt"
)

var outputPtr = flag.String("output", "", "a string")

// // check if the flag is valid and return the file name
func IsValidFlag(myFlag string) (bool, string) {
	isValid := true
	// // hide the first line
	// flag.CommandLine.SetOutput(io.Discard)
	// // catch if there is an error
	// flag.Usage = func() {
	// 	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")

	// 	// remove the last line
	// 	os.Exit(0)
	// }
	// extract the flag
	// flag.Parse()
	fmt.Println("myFlag", myFlag)
	// fileExt := path.Ext(myFlag)

	outputFile := *outputPtr
	if len(myFlag) >= 9 && myFlag[:9] != "--output=" {
		isValid = false
	}
	// if len(outputFile) >= 1 {
	// 	isValid = true
	// }
	return isValid, outputFile
}
