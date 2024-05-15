package handleFlag

import (
	"fmt"
	"os"
)

// make sure only one flag is used
// return that flag and it's value
func IsValidFlag(myFlags []string) (bool, string, bool, string) {
	isOutput := false
	isColor := false

	outputFile := ""
	color := ""

	isOutput, outputFile = checkIfOutput(myFlags, 0)

	if !isOutput && len(myFlags) >= 2 {
		isOutput, outputFile = checkIfOutput(myFlags, 1)
	}

	/// check if the first is --color
	isColor, color = checkIfColor(myFlags, 0)
	/// if the first is not --color
	if !isColor && len(myFlags) >= 2 {
		isColor, color = checkIfColor(myFlags, 1)
	}
	fmt.Println("flags : ", isOutput, isColor)
	if isOutput && isColor {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(1)
	}
	return isOutput, outputFile, isColor, color
}

func checkIfOutput(myFlags []string, argIndex int) (bool, string) {
	output := false
	outputFile := ""
	arg := myFlags[argIndex]

	// handle out of range
	if len(arg) >= 9 && arg[:9] == "--output=" {
		outputFile = arg[9:]
		output = true
	}
	return output, outputFile
}

func checkIfColor(myFlags []string, argIndex int) (bool, string) {
	isColor := false
	color := ""
	arg := myFlags[argIndex]
	if len(arg) >= 8 && arg[:8] == "--color=" {
		color = arg[8:]
		isColor = true
	}
	return isColor, color
}

func getColorANSI(color string) string {
	colors := map[string]string{"\033[0m": "Reset", "\033[31m": "Red", "\033[32m": "Green", "\033[33m": "Yellow", "Blue": "\033[34m", "Magenta": "\033[35m", "Cyan": "\033[36m", "Gray": "\033[37m", "White": "\033[97m"}
	fmt.Println(colors[color])
	return colors[color]
}
