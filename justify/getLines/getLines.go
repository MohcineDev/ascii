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
	bannerIndex := 0
	//	fmt.Println("GetLines args :", args, len(args))
	///if the flag is valid IsOutput = true
	if handleFlag.IsOutput {
		if len(args) == 1 { // args = [flag, text ]
		} else if len(args) == 2 {
			inputIndex = 1
		} else if len(args) == 3 {
			bannerIndex = 2
			inputIndex = 1
			bannerFile = "./Banners/" + args[bannerIndex]
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
	} else if handleFlag.IsAlign {
		////if there is a color flag
		if len(args) == 2 {
			inputIndex = 1
		} else if len(args) == 3 {
			inputIndex = 2
		}
	} else {
		if len(args) == 1 {
		} else if len(args) == 2 {
			bannerIndex = 1
			/// fs projLettersIndex := ect // no flag
			bannerFile = "./Banners/" + args[bannerIndex]
		} else if len(args) >= 3 && string(args[0][0]) != "-" {

			// fs usage message
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			os.Exit(2)
			// remove the last line
		}
	}
	file, err := os.ReadFile(bannerFile)
	fmt.Println("bannerFile : ", bannerFile)
	if err != nil {
		fmt.Println("aError :", args[bannerIndex], "file not found")

		return []string{}, ""
	}
	lines := strings.Split(string(file), "\n")

	return lines, args[inputIndex]
}

func GetLettersIndex(input string, letters string) []int {
	var indexes []int
	///replace letters to color with one char for easier search
	mm := strings.ReplaceAll(input, letters, "é")
	myInput := []rune(mm)

	for i, v := range myInput {
		if v == 'é' {
			indexes = append(indexes, i)
		}
	}

	///if there is an index
	if len(indexes) > 0 {
		for i := 1; i < len(indexes); i++ {
			indexes[i] += len(letters) - 1
			fmt.Println("i : ", indexes[i])
		}

		a := len(indexes)

		for i := 0; i < a; i++ {
			for j := 1; j < len(letters); j++ {
				indexes = append(indexes, indexes[i]+j)
			}
		}
	}

	return indexes
}
