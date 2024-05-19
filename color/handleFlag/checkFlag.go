package handleFlag

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	isOutput = false
	isColor  = false

	outputFile = ""
	Color      = ""
)

// make sure only one flag is used
// return that flag and it's value
func IsValidFlag(myFlags []string) (bool, string, bool, string) {
	isOutput, outputFile = checkIfOutput(myFlags, 0)

	if !isOutput && len(myFlags) >= 2 {
		isOutput, outputFile = checkIfOutput(myFlags, 1)
	}

	/// check if the first is --color
	isColor, Color = checkIfColor(myFlags, 0)
	/// if the first is not --color
	if !isColor && len(myFlags) >= 2 {
		isColor, Color = checkIfColor(myFlags, 1)
	}
	if isOutput && isColor {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(1)
	}
	return isOutput, outputFile, isColor, Color
}

func checkIfOutput(myFlags []string, argIndex int) (bool, string) {
	output := false
	outputFile := ""
	arg := myFlags[argIndex]

	// handle out of range
	if len(arg) >= 9 && arg[:9] == "--output=" {
		// if len(arg) >= 9 && arg[:9] == "--output=" || len(arg) >= 8 && arg[:8] == "--output" {
		if len(arg) >= 9 {
			outputFile = arg[9:]
		}

		output = true
	}
	return output, outputFile
}

/*
check if there is only one dash
*/
func checkIfColor(myFlags []string, argIndex int) (bool, string) {
	isColor := false
	color := ""
	arg := myFlags[argIndex]
	if len(arg) >= 8 && arg[:8] == "--color=" {
		if len(arg) >= 8 {
			color = arg[8:]
		}

		isColor = true
	}

	if isColor {
		if strings.Contains(color, "(") {
			// color = "\033[38;2;138;150;240m"
			color = getRgbColor(color)
		} else {
			color = getANSIColor(strings.ToLower(color))
		}
	}

	return isColor, color
}

func getANSIColor(color string) string {
	colors := map[string]string{
		"red": "\033[31m", "green": "\033[32m", "yellow": "\033[33m", "blue": "\033[34m",
		"magenta": "\033[35m", "cyan": "\033[36m", "gray": "\033[37m", "white": "\033[97m", "bright black": "\033[90m", "bright Red": "\033[91m",
		"bright Green": "\033[92m", "bright Yellow": "\033[93m", "bright blue": "\033[94m", "bright Magenta": "\033[95m", "bright Cyan": "\033[96m",
	}
	return colors[color]
}

func getRgbColor(rgbInput string) string {
	rgb := ""

	if strings.ToLower(rgbInput[:4]) != "rgb(" || string(rgbInput[len(rgbInput)-1:]) != ")" {
		fmt.Println("ERROR : COLOR NOT FOUND!!")
		os.Exit(1)
	}
	r, _ := regexp.Compile(`[0-9]+,[0-9]+,[0-9]+`)
	onlyNbr := r.FindString(rgbInput)

	if len(onlyNbr) >= 1 {
		rgb = "\033[38;2;" + strings.ReplaceAll(onlyNbr, ",", ";") + "m"
	}

	return rgb
}

func getColor(color string) {
	fmt.Println("getColor : ", color)
	myUrl := "https://csscolorsapi.com/api/colors/" + color
	// myUrl := "https://www.thecolorapi.com/scheme?rgb=0,71,171"
	res, err := http.Get(myUrl)
	if err != nil {
		fmt.Println("err")
	}

	content, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err")
		os.Exit(1)
	}
	fmt.Println("res:", string(content))
}
