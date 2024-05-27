package getLines

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"example.moh/handleFlag"
)

func MakeArt(input string, lines []string) (result string) {
	words := strings.Split(input, "\\n")
	newLineCount := strings.Count(input, "\\n")
	count := 0
	var lettersIndex []int
	if len(input) == 0 {
		return
	}
	/// to display correctly in the file
	// result = append(result, "")
	for a := 0; a < len(words); a++ {
		if len(LettersToColor) >= 1 {
			lettersIndex = Index(words[a], LettersToColor)
		}
		for i := 1; i < 9; i++ {

			for index, char := range words[a] {
				if int(char) < 32 || int(char) > 126 {
					fmt.Println("Error : char '", string(char), "' not found!!")
					// return
					os.Exit(1)
				}
				s := (int(char) - 32) * 9

				asciiLine := lines[s+i]
				///for the third file
				// asciiLine = strings.ReplaceAll(asciiLine, "\r", "")

				if slices.Contains(lettersIndex, index) || !LettersProvided {
					// result = append(result, handleFlag.Color+asciiLine+"\033[0m")
					result += handleFlag.Color + asciiLine + "\033[0m"
				} else {
					// result = append(result, asciiLine)
					result += asciiLine
				}

			}
			// result = append(result, "\n")
			if words[a] != "" {
				result += "\n"
			}

		}

		if count < newLineCount && words[a] == "" {
			// result = append(result, "\n")
			result += "\n"
			count++
		}
	}
	return result
}
