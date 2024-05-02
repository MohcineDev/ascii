package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arg1 := os.Args[1:]

	if len(arg1) < 1 {
		fmt.Println("Error please enter an input")
		return
	}
	/// hanlde file extension and args length

	// file, err := os.ReadFile("standard.txt")
	file, err := os.ReadFile("thinkertoy.txt")
	if err != nil {
		fmt.Println("Error standard.txt not found")
		return
	}
	line := strings.Split(string(file), "\n")

	words := []string{}

	// /split the first argument with line break
	if arg1[0] != "\\n" {
		words = strings.Split(arg1[0], "\\n")
	} 
	var result []string
	isLine := false

	if len(arg1[0]) == 0 {
		return
	}
	for a := 0; a < len(words); a++ {
		for i := 1; i < 9; i++ {
			isLine = false

			for _, char := range words[a] {
				s := (int(char) - 32) * 9

				if s > 856 {

					fmt.Println("Error : Your input is not found.!!")
					return
				}
				asciiLine := line[s+i]
				asciiLine = strings.ReplaceAll(asciiLine, "\r", "")
				result = append(result, asciiLine)
				isLine = true
			}
			if isLine {
				result = append(result, "\n")
			}
		}

		if len(words[a]) == 0 {
			result = append(result, "\n")
		}

	}

	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
	}
}
