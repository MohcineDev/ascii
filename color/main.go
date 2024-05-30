package main

import (
	"fmt"
	"os"

	"example.moh/getLines"
	"example.moh/handleFlag"
)

func main() {
	var result string

	lines, input := getLines.GetLines()
	result = getLines.MakeArt(input, lines)

	if !handleFlag.IsOutput {
		fmt.Print(result)
		return
	} 
	writingErr := os.WriteFile(handleFlag.OutputFile, []byte(result), 0o644)
	////IF THERE IS AN ERROR WRITING THE FILE! EX :
	if writingErr != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	}

}
