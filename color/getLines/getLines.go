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
	///handle args error
	argsError, args := handleArgs.CheckArgs(os.Args[1:])
	if argsError != nil {
		fmt.Println(argsError)
		os.Exit(0) /// it stops the test
	}

	bannerFile := "./Banners/standard.txt"
	inputIndex := 0

	///if the flag is valid IsOutput = true
	if handleFlag.IsOutput {
		if len(args) == 2 {
			inputIndex = 1
		} else if len(args) == 3 {
			inputIndex = 1
			bannerFile = "./Banners/" + args[2]
		}
	} else if handleFlag.IsColor {
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
			bannerFile = "./Banners/" + args[1]
		} else if len(args) >= 3 && string(args[0][0]) != "-" {
			// fs usage message
			fmt.Println("wwwwwUsage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			os.Exit(0)
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
