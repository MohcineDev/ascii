package getLines

import (
	"fmt"
	"os"
	"strings"

	"example.moh/handleArgs"
)

func GetLines() ([]string, string) {
	myArgs := os.Args[1:]

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
	if args[len(args)-1] == "validFlag" {
		if len(args) == 2 { // args = [flag, text, "validFlag"]
			bannerFile = "../standard.txt"
		} else if len(args) == 3 {
			fmt.Println("len 3")
			inputIndex = 1
			bannerFile = "./Banners/standard.txt"
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
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			// remove the last line
			os.Exit(0)
		}
	}
	fmt.Println(bannerFile)
	fmt.Println(args[bannerIndex])
	file, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error :", bannerFile, "file not found")
		return []string{}, ""

	}
	lines := strings.Split(string(file), "\n")

	return lines, args[inputIndex]
}
