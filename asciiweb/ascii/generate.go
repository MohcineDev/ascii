package ascii

import (
	"fmt"
	"os"
	"strings"
)

func Generate(input, banner string) string {
	lines := GetLines(banner)
	fmt.Println("from generate")
	var result []string
	count := 0
	resultS := ""
	input = strings.ReplaceAll(input, `\r\n`, "\r\n")

	words := strings.Split(input, "\r\n")
	newLineCount := strings.Count(input, "\n")

	if len(input) == 0 {
		return ""
		// return []string{}
	}
	for a := 0; a < len(words); a++ {
		for i := 1; i < 9; i++ {
			for _, char := range words[a] {

				if int(char) < 32 || int(char) > 126 {
					fmt.Printf("Error : char '%v' not found!!", string(char))
					os.Exit(1)
				}
				s := (int(char) - 32) * 9

				asciiLine := lines[s+i]
				///for the third file
				//asciiLine = strings.ReplaceAll(asciiLine, "\r", "")

				resultS += string(asciiLine)
				result = append(result, asciiLine)

			}
			resultS += "\n"
			result = append(result, "\n")

		}

		if count < newLineCount && words[a] == "" {
			result = append(result, "\n")
			count++
		}

	}
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
	}

	fmt.Println("len : ", len(result))
	// fmt.Println(result)
	fmt.Println(resultS)

	return resultS
}
