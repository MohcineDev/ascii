package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"example/handleArgs"
)

/// return the lines of the selected ex : (standard , shadow...) file
///and the first arg

func getLines() ([]string, string, error) {
	myargs := os.Args[1:]

	///handle args errors
	argsError, args := handleArgs.CheckArgs(myargs)
	if argsError != nil {

		fmt.Println(argsError)
		//	os.Exit(1) /// it stops the test
		return []string{}, "", errors.New("err")
	}

	bannerFile := ""
	if len(args) == 1 {
		bannerFile = "../standard.txt"
	} else if len(args) == 2 {
		bannerFile = "../" + args[1] + ".txt"
	}

	file, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error ", args[1], " not found")
		return []string{}, "", errors.New("err")

	}
	lines := strings.Split(string(file), "\n")

	return lines, args[0], nil
}
