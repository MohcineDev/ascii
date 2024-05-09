package main

import (
	"fmt"
	"getLines"
	"strings"
)

func main() {
	lines, input, _ := getLines.GetLines()
	/// hanlde file extension

	// /split the first argument with line break
	words := strings.Split(input, "\\n")
	newLineCount := strings.Count(input, "\\n")

	var result []string
	endLine := false
	count := 0

	if len(input) == 0 {
		return
	}

	for a := 0; a < len(words); a++ {
		for i := 1; i < 9; i++ {
			endLine = false

			for _, char := range words[a] {
				if int(char) < 32 || int(char) > 126 {
					fmt.Println("Error : char '", string(char), "' not found!!")
					return
				}
				s := (int(char) - 32) * 9

				asciiLine := lines[s+i]
				///for the third file
				asciiLine = strings.ReplaceAll(asciiLine, "\r", "")
				result = append(result, asciiLine)
				endLine = true
			}
			if endLine {
				result = append(result, "\n")
			}
		}

		if count < newLineCount && words[a] == "" {
			result = append(result, "\n")
			count++
		}

	}

	////Print result
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
	}
}
