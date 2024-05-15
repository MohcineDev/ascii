package handleFlag

import (
	"fmt"
	"os"
)

var (
	isOutput = false
	isColor  = false

	outputFile = ""
	color      = ""
)

// make sure only one flag is used
// return that flag and it's value
func IsValidFlag(myFlags []string) (bool, string, bool, string) {
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
	if len(arg) >= 8 && arg[:8] == "--output" {
		// if len(arg) >= 9 && arg[:9] == "--output=" || len(arg) >= 8 && arg[:8] == "--output" {
		if len(arg) >= 9 {
			outputFile = arg[9:]
		}
		output = true
	}
	return output, outputFile
}

func checkIfColor(myFlags []string, argIndex int) (bool, string) {
	isColor := false
	color := ""
	arg := myFlags[argIndex]
	if len(arg) >= 7 && arg[:7] == "--color" {
		if len(arg) >= 8 {
			color = arg[8:]
		}
		isColor = true
	}
	return isColor, color
}

// /return the color
func GetColor() string {
	return getColorANSI(color)
}

func getColorANSI(color string) string {
	colors := map[string]string{
		"Reset": "\033[0m", "red": "\033[31m",
		"green": "\033[32m", "yellow": "\033[33m", "blue": "\033[34m",
		"magenta": "\033[35m", "cyan": "\033[36m", "gray": "\033[37m", "white": "\033[97m",
	}
	fmt.Println(colors[color])
	return colors[color]
}
