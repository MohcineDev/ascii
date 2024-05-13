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
	// hide the first line in the terminal
	flag.CommandLine.SetOutput(io.Discard)
	// catch if there is an error
	flag.Usage = func() {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		// remove the last line
	}
	///Parse parses the command-line flags from os.Args[1:]. Must be called after all flags are defined and before flags are accessed by the program.
	flag.Parse()
	// save args without the flag
	myargs := flag.Args()

	///handle args error
	argsError, args := handleArgs.CheckArgs(myargs)
	if argsError != nil {
		fmt.Println(argsError)
		os.Exit(0) /// it stops the test
	}
	bannerFile := ""
	///if the flag is valid add "validFlag" to the end of the args slice on the checkArgs function
	if args[len(args)-1] == "validFlag" {
		if len(args) == 2 { // args = [flag, text, "validFlag"]
			bannerFile = "../standard.txt"
		} else if len(args) == 3 {
			bannerFile = "../" + args[1]
		}
	} else {
		if len(args) == 1 {
			bannerFile = "../standard.txt"
		} else if len(args) == 2 {
			/// fs project // no flag
			bannerFile = "../" + args[1]
		}
	}

	file, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error :", args[1], "file not found")
		return []string{}, ""

	}
	lines := strings.Split(string(file), "\n")

	return lines, args[0]
}
