package handleFlag

import (
	"flag"
	"os"
)

var outputPtr = flag.String("output", "", "a string")

// // check if the flag is valid and return the file name
func IsValidFlag() (bool, string) {
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
	flag.Parse()

	myFlag := os.Args[1]
	//	fileExt := path.Ext(os.Args[1])

	outputFile := *outputPtr
	if len(myFlag) >= 9 && myFlag[:9] != "--output=" || flag.NFlag() < 1 {
		isValid = false
	}
	// if len(outputFile) >= 1 {
	// 	isValid = true
	// }
	return isValid, outputFile
}
