package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"example.moh/getLines"
	"example.moh/handleFlag"
)

//TODO
/// when using output flag display output usage msg
/// when using color flag display color usage msg
///

func main() {
	var result []string
	endLine := false
	count := 0
	var lettersIndex []int

	lines, input := getLines.GetLines()

	words := strings.Split(input, "\\n")
	newLineCount := strings.Count(input, "\\n")

	if len(input) == 0 {
		return
	}
	/// to display correctly in the file
	result = append(result, "")
	letterIndex := 0
	for a := 0; a < len(words); a++ {
		if len(getLines.LettersToColor) >= 1 {
			lettersIndex = getLines.GetLettersIndex(words[a], getLines.LettersToColor)
		}
		for i := 1; i < 9; i++ {
			endLine = false
			letterIndex = 0

			for _, char := range words[a] {
				if int(char) < 32 || int(char) > 126 {
					fmt.Println("Error : char '", string(char), "' not found!!")
					// return
					os.Exit(1)
				}
				s := (int(char) - 32) * 9

				asciiLine := lines[s+i]
				///for the third file
				asciiLine = strings.ReplaceAll(asciiLine, "\r", "")

				if slices.Contains(lettersIndex, letterIndex) || !getLines.LettersProvided {
					result = append(result, handleFlag.Color+asciiLine+"\033[0m")
					// result = append(result, handleFlag.GetColor()+asciiLine+"\033[0m")
				} else {
					result = append(result, asciiLine)
				}

				endLine = true

				letterIndex++
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
	//////////////// O U T P U T ///////////////////
	isOutput := false
	isColor := false

	outputFile := ""
	color := ""
	isOutput, outputFile, isColor, color = handleFlag.IsValidFlag(os.Args[1:])

	if !isOutput {
		// chosenColor := "\033[31m"
		// print result
		for i := 0; i < len(result); i++ {
			fmt.Print(result[i])
		}

		if isColor {
			fmt.Println(color)
		}
	} else if len(os.Args[1:]) >= 2 && isOutput {
		writingErr := os.WriteFile(outputFile, []byte(strings.Join(result, "")), 0o644)
		////IF THERE IS AN ERROR WRITING THE FILE! EX :
		if writingErr != nil {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		}
	}
}
