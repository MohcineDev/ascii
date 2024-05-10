package handleFlag

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var outputPtr = flag.String("output", "", "a string")

// // check if the flag is valid and return the file name
func IsValidFlag() (bool, string) {
	isValid := false
	// hide the first line
	flag.CommandLine.SetOutput(io.Discard)
	// catch if there is an error
	flag.Usage = func() {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")

		// remove the last line
		os.Exit(0)
	}
	// extract the flag
	flag.Parse()
	outputFile := *outputPtr
	if len(outputFile) >= 1 {
		isValid = true
	}
	fmt.Println("------ valid Flag ", isValid)
	return isValid, outputFile
}
