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

	fmt.Printf("is align  : %v , alignment : %s", handleFlag.IsAlign, handleFlag.Alignment)
	fmt.Println("")
	/// to display correctly in the file
	result = append(result, "")
	letterIndex := 0
	//get word length using style
	for a := 0; a < len(words); a++ {
		lineWidth := getLineWidth(words[a], lines)
		if len(getLines.LettersToColor) >= 1 {
			lettersIndex = getLines.GetLettersIndex(words[a], getLines.LettersToColor)
		}
		wordsBySpace := []string{}
		if handleFlag.IsAlign && handleFlag.Alignment == "justify" {

			wordsBySpace = strings.Split(words[a], " ")
			fmt.Println(wordsBySpace)
		} else {
			wordsBySpace = []string{words[a]}
		}
		// if len(wordsBySpace) > 1 {

		// }
		for i := 1; i < 9; i++ {
			if handleFlag.IsAlign && handleFlag.Alignment == "right" {

				printSpaces(terminalWidth - lineWidth)
			}
			///center
			if handleFlag.IsAlign && handleFlag.Alignment == "center" {

				printSpaces((terminalWidth - lineWidth) / 2)
			}

			endLine = false
			letterIndex = 0
			for j := 0; j < len(wordsBySpace); j++ {
				if handleFlag.IsAlign && handleFlag.Alignment == "justify" {
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
					if handleFlag.IsAlign {
						fmt.Print(asciiLine)
					}

					endLine = true
					letterIndex++
				}
			}
			if handleFlag.IsAlign {
				fmt.Print("\n")
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
	/*
	                     mm
	            mm
	   mm      mmm       mm
	   mm   mm    mmm    mm
	   terminal size / len(word / " ")-1 =  3
	   mm                  */

	// isOutput, outputFile, isColor, color = handleFlag.IsValidFlag(os.Args[1:])
	if !handleFlag.IsOutput && !handleFlag.IsAlign {
		fmt.Println("1312")
		// print result
		for i := 0; i < len(result); i++ {
			fmt.Print(result[i])
		}

	} else if len(os.Args[1:]) >= 2 && handleFlag.IsOutput {
		writingErr := os.WriteFile(handleFlag.OutputFile, []byte(strings.Join(result, "")), 0o644)
		////IF THERE IS AN ERROR WRITING THE FILE! EX :
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
		fmt.Print("-")
	}
}

func getLineWidth(word string, lines []string) int {
	i := 1
	lineWidth := 0
	if true {
		word = strings.ReplaceAll(word, " ", "")
	}
	for _, char := range word {

		s := (int(char) - 32) * 9

		lineWidth += len(lines[s+i])
		i++
		if i == 9 {
			i = 1
		}
	}
	return lineWidth
}

/*

func main() {
	var result []string
	endLine := false
	count := 0
	var lettersIndex []int
	///line width to use for justify project
	lineWidth := 0

	getTerminalWidth()
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
			if handleFlag.IsAlign {
				printSpaces(getTerminalWidth() - lineWidth)
			}
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
				//collect letters length for justify project
				if i == 1 {

					lineWidth += len(lines[s+3])
				}

				if slices.Contains(lettersIndex, letterIndex) || !getLines.LettersProvided {
					result = append(result, handleFlag.Color+asciiLine+"\033[0m")
					// result = append(result, handleFlag.GetColor()+asciiLine+"\033[0m")
				} else {
					if handleFlag.IsAlign {

						fmt.Print(asciiLine)
					} else {

						result = append(result, asciiLine)
					}
				}

				endLine = true
				letterIndex++
			}
			if handleFlag.IsAlign {

				fmt.Print("\n")
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

	// isOutput, outputFile, isColor, color = handleFlag.IsValidFlag(os.Args[1:])
	if !handleFlag.IsOutput && !handleFlag.IsAlign {
		// print result
		for i := 0; i < len(result); i++ {

			//1 of the last new line
			if result[i] != "\n" {
				printSpaces(getTerminalWidth() - lineWidth)

			}
			fmt.Print(result[i])
		}

	} else if len(os.Args[1:]) >= 2 && handleFlag.IsOutput {
		writingErr := os.WriteFile(handleFlag.OutputFile, []byte(strings.Join(result, "")), 0o644)
		////IF THERE IS AN ERROR WRITING THE FILE! EX :
		if writingErr != nil {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		}
	}
}

*/
