package getLines

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"example/handleArgs"
)

/// return the lines of the selected file ex : (standard , shadow...) file
///and the first arg

func GetLines() ([]string, string, error) {
	myargs := os.Args[1:]
	///index in the returned args slice
	inputIndex := 0
	bannerIndex := 0
	///handle args errors
	argsError, args := handleArgs.CheckArgs(myargs)
	if argsError != nil {

		fmt.Println(argsError)
		//	os.Exit(1) /// it stops the test
		return []string{}, "", errors.New("err")
	}

	bannerFile := ""
	if len(args) == 3 && args[2] == "validFlag" {
		bannerFile = "../standard.txt"
		inputIndex = 1
	} else if len(args) == 4 && args[3] == "validFlag" {
		bannerIndex = 2
		bannerFile = "../" + args[bannerIndex] + ".txt"
		inputIndex = 2 
	} else {
		if len(args) == 1 {
			bannerFile = "../standard.txt"
		} else if len(args) == 2 {
			bannerFile = "../" + args[bannerIndex] + ".txt"
		}
	}

	file, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error ", args[bannerIndex], " not found")
		return []string{}, "", errors.New("err")

	}
	lines := strings.Split(string(file), "\n")

	return lines, args[inputIndex], nil
}
