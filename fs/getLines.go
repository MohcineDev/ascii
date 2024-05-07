package main

import (
	"fmt"
	"os"
	"strings"

	"example/handleArgs"
)

func getLines() ([]string, string) {
	myargs := os.Args[1:]

	argsError, args := handleArgs.CheckArgs(myargs)
	if argsError != nil {
		fmt.Println(argsError)
		os.Exit(1)
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
		os.Exit(1)
	}
	lines := strings.Split(string(file), "\n")

	return lines, args[0]
}
