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
	file, _ := os.ReadFile("standard.txt")
	line := strings.Split(string(file), "\n")
	arg1 := os.Args[1]

	var newarr []string
	for i := 0; i < 9; i++ {
		for _, letter := range arg1 {
			s := (int(letter)-32)*9 + 1
			r := line[s+i]
			newarr = append(newarr, r)

		}
		newarr = append(newarr, "\n")

	}
	for i := 0; i < len(newarr); i++ {
		
		fmt.Print(newarr[i])
	}
}
