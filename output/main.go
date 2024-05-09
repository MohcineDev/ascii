package main

import (
	"fmt"
	"getLines"
	"handleFlag"
	"os"
	"strings"
)

func main() {
	var result []string
	endLine := false
	count := 0

	// /open asciii
	lines, input, _ := getLines.GetLines()

	words := strings.Split(input, "\\n")
	newLineCount := strings.Count(input, "\\n")

	if len(input) == 0 {
		return
	}
	/// to display correctly in the file
	result = append(result, "")
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

	///* end ascii

	// print result
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
	}

	//////////////// O U T P U T ///////////////////
	if len(os.Args[1:]) >= 2 {

		_, fileName := handleFlag.IsValidFlag()
		writingErr := os.WriteFile(fileName, []byte(strings.Join(result, " ")), 0o644)
		////IF THERE IS ANN ERROR WRITING THE FILE! EX :
		if writingErr != nil {
			fmt.Println("Error : can't write the file")
		}
	}
}
