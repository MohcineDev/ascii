package getLines

import (
	"fmt"
	"os"
	"strings"

	"example.moh/handleArgs"
)

func GetLines() ([]string, string, []int) {
	myArgs := os.Args[1:]
	LettersIndex := []int{}
	///handle args error
	argsError, args := handleArgs.CheckArgs(myArgs)
	if argsError != nil {
		fmt.Println(argsError)
		os.Exit(0) /// it stops the test
	}

	fmt.Println("getline args : ", args, len(args))
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
		////if there is a color flag
		bannerFile = "./Banners/standard.txt"
		if len(args) == 3 {
			inputIndex = 1
		} else if len(args) == 4 {
			inputIndex = 2
			LettersIndex = getLettersIndex(args[2], args[1])
		}
	} else {
		if len(args) == 1 {
			bannerFile = "../standard.txt"
		} else if len(args) == 2 {
			bannerIndex = 1
			/// fs projLettersIndex := ect // no flag
			bannerFile = "../" + args[bannerIndex]
		} else if len(args) >= 3 {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			// remove the last line
			os.Exit(0)
		}
	}
	file, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error :", bannerFile, "file not found")
		return []string{}, "", []int{}

	}
	lines := strings.Split(string(file), "\n")

	return lines, args[inputIndex], LettersIndex
}

func getLettersIndex(input string, letters string) []int {
	var indexes []int
	mm := strings.ReplaceAll(input, letters, "ب")
	myInput := []rune(mm)

	fmt.Println(mm)
	for i, v := range myInput {
		if v == 'ب' {
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
