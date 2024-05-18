package getLines

import (
	"fmt"
	"os"
	"strings"

	"example.moh/handleArgs"
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

	///if the flag is valid add "validFlag" to the end of the args slice on the checkArgs function
	if args[len(args)-1] == "validFlag" {
		if len(args) == 2 { // args = [flag, text, "validFlag"]
			bannerFile = "../standard.txt"
		} else if len(args) == 3 {
			inputIndex = 1
			bannerFile = "./Banners/standard.txt"
		} else if len(args) == 4 {
			bannerIndex = 2
			inputIndex = 1
			bannerFile = "../" + args[bannerIndex]
		}
	} else if args[len(args)-1] == "colorFlag" {

		////
		////if there is a color flag
		bannerFile = "./Banners/standard.txt"
		if len(args) == 3 {
			///no letterstocolor provided
			LettersProvided = false
			inputIndex = 1
		} else if len(args) == 4 {
			LettersToColor = args[1]
			inputIndex = 2
		}
	} else {
		if len(args) == 1 {
			bannerFile = "../standard.txt"
		} else if len(args) == 2 {
			bannerIndex = 1
			/// fs projLettersIndex := ect // no flag
			bannerFile = "../" + args[bannerIndex]
		} else if len(args) >= 3 && string(args[0][0]) != "-" {

			fmt.Println("args geline : ", args)
			// fs usage message
			fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
			os.Exit(2)
			// remove the last line
		}
	}
	file, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error :", bannerFile, "file not found")

		return []string{}, ""
	}
	lines := strings.Split(string(file), "\n")

	return lines, args[inputIndex]
}

func GetLettersIndex(input string, letters string) []int {
	var indexes []int
	mm := strings.ReplaceAll(input, letters, "é")
	myInput := []rune(mm)

	for i, v := range myInput {
		if v == 'é' {
			indexes = append(indexes, i)
		}
	}
	if len(indexes) > 1 {
		for i := 1; i < len(indexes); i++ {
			indexes[i] += len(letters) - 1
			fmt.Println("i : ", indexes[i])
		}
	}

	a := len(indexes)
	for i := 0; i < a; i++ {
		for j := 1; j < len(letters); j++ {
			indexes = append(indexes, indexes[i]+j)
		}
	}

	return indexes
}
