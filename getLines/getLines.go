package getLines

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"example.moh/handleArgs"
)

/// return the lines of the selected file ex : (standard , shadow...) file
/// and the first arg

func GetLines() ([]string, string) {
	myArgs := os.Args[1:]
	// hide the first line in the terminal
	flag.CommandLine.SetOutput(io.Discard)
	// catch if there is an error
	// flag.Usage = func() {
	// 	fmt.Println("from usage \n Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	// 	// remove the last line
	// 	os.Exit(0)
	// }
	///Parse parses the command-line flags from os.Args[1:]. Must be called after all flags are defined and before flags are accessed by the program.
	//flag.Parse()

	// save args without the flag
	//	myargs := flag.Args()

	///handle args error
	argsError, args := handleArgs.CheckArgs(myArgs)
	if argsError != nil {
		fmt.Println(argsError)
		os.Exit(0) /// it stops the test
	}
	bannerFile := ""
	inputIndex := 0
	bannerIndex := 0
	///if the flag is valid add "validFlag" to the end of the args slice on the checkArgs function
	fmt.Println("getline args: ", args)
	if args[len(args)-1] == "validFlag" {
		if len(args) == 2 { // args = [flag, text, "validFlag"]
			bannerFile = "../standard.txt"
		} else if len(args) == 3 {
			fmt.Println("44")
			inputIndex = 1
			bannerFile = "../standard.txt"
		} else if len(args) == 4 {
			bannerIndex = 2
			inputIndex = 1
			bannerFile = "../" + args[bannerIndex]
		}
	} else {
		if len(args) == 1 {
			bannerFile = "../standard.txt"
		} else if len(args) == 2 {
			bannerIndex = 1
			/// fs project // no flag
			bannerFile = "../" + args[bannerIndex]
		} else if len(args) >= 3 {
			fmt.Println("63:Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			// remove the last line
			os.Exit(0)
		}
	}
	fmt.Println("68:", bannerFile)
	file, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error :", args[bannerIndex], "file not found")
		return []string{}, ""

	}
	lines := strings.Split(string(file), "\n")

	return lines, args[inputIndex]
}
