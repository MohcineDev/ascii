package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// if len(os.Args) == 1 {
	// 	return
	// }
	// //Writing arguments in a single string
	// str := os.Args[1]
	// for _, v := range os.Args[2:] {
	// 	str += " " + v
	// }
	// //2. Checking weather str contain "\n" or not ---> executing the ascii-art
	// severallines := false
	// for i, v := range str {
	// 	if v == 'n' && i-1 == '\\' {
	// 		severallines = true
	// break
	// 	}
	// }
	// //3. Writing text line by line into res
	// res := ""
	// if severallines {
	// 	args := strings.Split(str, "\\n")
	// 	for _, word := range args {
	// 		for i := 0; i < 8; i++ {
	// 			for _, letter := range word {
	// 				res += GetLine(1 + int(letter-' ')*9 + i)
	// 			}
	// 			fmt.Println(res)
	// 			res = ""
	// 		}
	// 	}
	// } else {
	// 	for i := 0; i < 8; i++ {
	// 		for _, letter := range str {
	// 			res += GetLine(1 + int(letter-' ')*9 + i)
	// 			fmt.Println("res : ", res)
	// 		}
	// 		res = ""
	// 	}
	// }
	lines()
}

func lines() {
	arg1 := os.Args[1:]

	if len(arg1) < 1 {
		fmt.Println("Error please enter an input")
		return
	}

	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error standard.txt not found")
		return
	}
	line := strings.Split(string(file), "\n")

	words := []string{}

	// /split the argument with line break
	if arg1[0] != "\\n" {
		words = strings.Split(arg1[0], "\\n")
	} else {
		fmt.Println("")
	}
	var result []string
	isLine := false

	/////
	if len(arg1[0]) == 0 {
		return
	}


	for a := 0; a < len(words); a++ { 
		for i := 1; i < 9; i++ {
			isLine = false

			for _, char := range words[a] {
				s := (int(char) - 32) * 9

				if s > 856 {

					fmt.Println("Error : Your input is not found.!!")
					return
				}
				asciiLine := line[s+i]

				result = append(result, asciiLine)
				isLine = true
			}
			if isLine {
				result = append(result, "\n")
			}
		}
		if len(words[a]) == 0 {
			result = append(result, "\n")
		}

	}
	/////print result
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i])
	}
}
