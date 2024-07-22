package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strconv"
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
	///line width to use for justify project

	terminalWidth := getTerminalWidth()
	lines, input := getLines.GetLines()
	words := strings.Split(input, "\\n")
	newLineCount := strings.Count(input, "\\n")

	if len(input) == 0 {
		return
	}

	/// to display correctly in the file
	result = append(result, "")
	letterIndex := 0
	//get word length using style
	for a := 0; a < len(words); a++ {
		lineWidth := getLineWidth(words[a], lines)
		if len(getLines.LettersToColor) >= 1 {
			lettersIndex = getLines.GetLettersIndex(words[a], getLines.LettersToColor)
		}
		var wordsBySpace = []string{}
		if handleFlag.Alignment == "justify" {
			//words[a] = strings.TrimSpace(words[a])
			wordsBySpace = strings.Split(words[a], " ")

		} else {
			wordsBySpace = []string{words[a]}
		}

		for i := 1; i < 9; i++ {
			if len(wordsBySpace[0]) > 1 {
				switch handleFlag.Alignment {
				case "right":
					printSpaces(terminalWidth - lineWidth)
				case "center":
					printSpaces((terminalWidth - lineWidth) / 2)
				}

			}
			endLine = false
			letterIndex = 0
			for j := 0; j < len(wordsBySpace); j++ {
				if handleFlag.Alignment == "justify" {
					if j > 0 {
						printSpaces((terminalWidth - lineWidth) / (len(wordsBySpace) - 1))
					}
				}
				for _, char := range wordsBySpace[j] {

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
					} else {

						result = append(result, asciiLine)
					}
					if handleFlag.IsAlign && wordsBySpace[j] != "" {
						fmt.Print(asciiLine)
					}
					endLine = true
					letterIndex++
				}
			}
			if handleFlag.IsAlign && endLine {
				//if the of line is reached and the align flag is present print a new line
				fmt.Print("\n")
			}

			//this used in other projects except [justify]
			if endLine {
				result = append(result, "\n")
			}
		}
		if count < newLineCount && handleFlag.IsAlign && words[a] == "" {
			fmt.Print("\n")
			count++
		}
		//used in all projects except [justify]
		//&& !handleFlag.IsAlign  to prevent this condition from execution
		if count < newLineCount && words[a] == "" && !handleFlag.IsAlign {
			result = append(result, "\n")
			count++
		}

	}
	//////////////// O U T P U T ///////////////////

	////[justify] is printed above and [output] is saved in the file
	if !handleFlag.IsOutput && !handleFlag.IsAlign {

		// print result
		for i := 0; i < len(result); i++ {
			fmt.Print(result[i])
		}

	} else if len(os.Args[1:]) >= 2 && handleFlag.IsOutput {
		writingErr := os.WriteFile(handleFlag.OutputFile, []byte(strings.Join(result, "")), 0o644)
		////IF THERE IS AN ERROR WRITING THE FILE!
		if writingErr != nil {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		}
	}
}

func getTerminalWidth() int {
	// cmd := exec.Command("stty", "size")
	cmd := exec.Command("tput", "cols")

	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	w, err := strconv.Atoi(string(out)[:len(string(out))-1])
	if err != nil {
		fmt.Println(err)
	}
	return w
}

func printSpaces(width int) {
	for i := 0; i < width; i++ {
		fmt.Print(" ")
	}
}

func getLineWidth(word string, lines []string) int {
	lineWidth := 0
	// fmt.Println(word)
	if handleFlag.Alignment == "justify" {
		///remove spaces to only count the chars
		word = strings.ReplaceAll(word, " ", "")
	}
	for _, char := range word {

		if int(char) < 32 || int(char) > 126 {
			fmt.Println("Error : char '", string(char), "' not found!!")
			// return
			os.Exit(1)
		}
		s := (int(char) - 32) * 9
		//add the length of the char based on the used style
		lineWidth = lineWidth + len(lines[s+1])
		// fmt.Println(lines[s+2], len(lines[s+1]))
		// fmt.Println("")
		// fmt.Println("")
	}
	fmt.Println(lineWidth)
	return lineWidth
}
