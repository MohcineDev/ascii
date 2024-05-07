package main

import (
	"fmt"
	"os"
	"strings"

	"example/handleArgs"
)

func main() {
	argsError, input := handleArgs.CheckArgs()
	if argsError != nil {
		fmt.Println(argsError)
		return
	}

	/// hanlde file extension

	file, err := os.ReadFile("../standard.txt")
	if err != nil {
		fmt.Println("Error standard.txt not found")
		return
	}
	line := strings.Split(string(file), "\n")

	words := []string{}

	// /split the first argument with line break
	words = strings.Split(input[0], "\\n")
	newLineCount := strings.Count(input[0], "\\n")

	var result []string
	isLine := false
	count := 0
	if len(input[0]) == 0 {
		return
	}
	for a := 0; a < len(words); a++ {
		for i := 1; i < 9; i++ {
			isLine = false

			for _, char := range words[a] {
				if int(char) < 32 || int(char) > 126 {
					fmt.Println("Error : char", char, " is not found.!!")
					return
				}
				s := (int(char) - 32) * 9

				asciiLine := line[s+i]
				///for the third file
				// asciiLine = strings.ReplaceAll(asciiLine, "\r", "")
				result = append(result, asciiLine)
				isLine = true
			}
			if isLine {
				result = append(result, "\n")
			}
		}

		if count < newLineCount && words[a] == "" {
			result = append(result, "\n")
			count++
		}

	}
//print result
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
	}
}
