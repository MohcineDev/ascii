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

	bannerFile := ""
	inputIndex := 0
	bannerIndex := 0
	//	fmt.Println("GetLines args :", args, len(args))
	///if the flag is valid IsOutput = true
	if handleFlag.IsOutput {
		if len(args) == 1 { // args = [flag, text ]
			bannerFile = "./Banners/standard.txt"
		} else if len(args) == 2 {
			inputIndex = 1
			bannerFile = "./Banners/standard.txt"
		} else if len(args) == 3 {
			bannerIndex = 2
			inputIndex = 1
			bannerFile = "../" + args[bannerIndex]
		}
	} else if handleFlag.IsColor {
		////if there is a color flag
		bannerFile = "./Banners/standard.txt"
		if len(args) == 2 {
			///no letterstocolor provided
			LettersProvided = false
			inputIndex = 1
		} else if len(args) == 3 {
			LettersToColor = args[1]
			inputIndex = 2
		}
	} else {
		if len(args) == 1 {
			bannerFile = "./Banners/standard.txt"
		} else if len(args) == 2 {
			bannerIndex = 1
			/// fs projLettersIndex := ect // no flag
			bannerFile = "../" + args[bannerIndex]
		} else if len(args) >= 3 && string(args[0][0]) != "-" {

			// fs usage message
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			os.Exit(2)
			// remove the last line
		}
	}
	file, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("aError :", bannerFile, "file not found")

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
	fmt.Println("indexes  : ", indexes)
	///if there is an index
	if len(indexes) > 0 {
		b := 0
		for i := 1; i < len(indexes); i++ {

			if i > 1 && len(letters) > 1 {
				indexes[i] += len(letters) - 1 + b

			} else {

				indexes[i] += len(letters) - 1
			}
			///3  8  14 18

			fmt.Println("i : ", indexes[i])
			b++
		}
		fmt.Println("new indexes : ", indexes)
		a := len(indexes)

		for i := 0; i < a; i++ {
			for j := 1; j < len(letters); j++ {
				indexes = append(indexes, indexes[i]+j)
			}
		}
	}
	fmt.Println("lqst indexes : ", indexes)

	return indexes
}

func Index(input, letters string) (slice []int) {
	for i := 0; i < len(input)-len(letters); i++ {
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
	fmt.Println(slice)
	return
}
