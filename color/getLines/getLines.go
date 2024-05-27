package getLines

import (
	"fmt"
	"os"
	"strings"

	"example.moh/handleArgs"
	"example.moh/handleFlag"
)

var (
	LettersToColor  string
	RGBColor        string
	LettersProvided bool = true
)

func GetLines() ([]string, string) {
	myArgs := os.Args[1:]

	///handle args error
	argsError, args := handleArgs.CheckArgs(myArgs)
	if argsError != nil {
		fmt.Println(argsError)
		os.Exit(0) /// it stops the test
	}

	bannerFile := "./Banners/standard.txt"
	inputIndex := 0
	fmt.Println("GetLines args :", args, len(args))
	///if the flag is valid IsOutput = true

	if handleFlag.IsOutput {
		fmt.Println(handleFlag.IsOutput)

		if len(args) == 1 { // args = [flag, text ]
		} else if len(args) == 2 {
			inputIndex = 1
		} else if len(args) == 3 {
			inputIndex = 1
			bannerFile = "./Banners/" + args[2]
		}
	} else if handleFlag.IsColor {
		bannerFile = "./Banners/standard.txt"
		////if there is a color flag
		if len(args) == 2 {
			///no letterstocolor provided
			LettersProvided = false
			inputIndex = 1
		} else if len(args) == 3 {
			LettersToColor = args[1]
			inputIndex = 2
		}
	} else {
		if len(args) == 2 {
			/// fs projLettersIndex := ect // no flag
			bannerFile = "./Banners/" + args[1]
		} else if len(args) >= 3 && string(args[0][0]) != "-" {

			// fs usage message
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			os.Exit(2)
			// remove the last line
		}
	}
	file, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("invalide Banner")

		return []string{}, ""
	}
	lines := strings.Split(string(file), "\n")

	return lines, args[inputIndex]
}

func Index(input, letters string) (slice []int) {
	for i := 0; i <= len(input)-len(letters); i++ {
		if letters == input[i:i+len(letters)] {
			slice = append(slice, i)
		}
	}
	sliceLength := len(slice)
	for i := 0; i < sliceLength; i++ {
		for j := 1; j < len(letters); j++ {
			slice = append(slice, slice[i]+j)
		}
	}
	return
}
