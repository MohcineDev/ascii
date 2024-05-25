package handleFlag

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var (
	IsOutput = false
	IsColor  = false
	IsAlign  = false

	OutputFile = ""
	Color      = ""
	Alignment  = ""
)

// make sure only one flag is used
// return that flag and it's value
func IsValidFlag(myFlags []string) (bool, string, bool, string) {
	IsOutput, OutputFile = checkIfOutput(myFlags, 0)
	if !IsOutput && len(myFlags) >= 2 {
		IsOutput, OutputFile = checkIfOutput(myFlags, 1)
	}
	/// check if the first is --color
	IsColor, Color = checkIfColor(myFlags, 0)
	/// if the first is not --color
	if !IsColor && len(myFlags) >= 2 {
		IsColor, Color = checkIfColor(myFlags, 1)
	}
	checkIfAlign(myFlags, 0)
	if !IsAlign && len(myFlags) >= 2 {
		checkIfAlign(myFlags, 1)
	}

	if IsOutput && IsColor {
		fmt.Println("0000000000000000000Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		os.Exit(1)
	}
	return IsOutput, OutputFile, IsColor, Color
}

func checkIfAlign(myFlags []string, argIndex int) {
	arg := myFlags[argIndex]
	// handle out of range
	if len(arg) >= 8 && arg[:8] == "--align=" {
		if len(arg) >= 8 {
			Alignment = arg[8:]
		}
		IsAlign = true
	}
}

func checkIfOutput(myFlags []string, argIndex int) (bool, string) {
	output := false
	OutputFile = ""
	arg := myFlags[argIndex]

	// handle out of range
	if len(arg) >= 9 && arg[:9] == "--output=" {
		// if len(arg) >= 9 && arg[:9] == "--output=" || len(arg) >= 8 && arg[:8] == "--output" {
		if len(arg) >= 9 {
			OutputFile = arg[9:]
		}

		output = true
	}
	return output, OutputFile
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
		if strings.Contains(color, "(") || strings.Contains(color, "#") {
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
		"brown": "\033[38;2;165;42;42m", "crimson": "\033[38;2;220;20;60m", "fuchsia": "\033[38;2;255;0;255m", "gold": "\033[38;2;255;215;0m", "indigo": "\033[38;2;75;0;130m", "lightblue": "\033[38;2;173;216;230m",
		"maroon": "\033[38;2;128;0;0m", "orange": "\033[38;2;255;165;0m", "pink": "\033[38;2;255;192;203m", "purple": "\033[38;2;128;0;128m", "salmon": "\033[38;2;250;128;114m",

		"black": "\033[38;2;0;0;0m", "silver": "\033[38;2;192;192;192m", "lime": "\033[38;2;0;255;0m", "navy": "\033[38;2;0;0;128m", "olive": "\033[38;2;128;128;0m", "teal": "\033[38;2;0;128;128m", "aqua": "\033[38;2;0;255;255m",
		"hot pink": "\033[38;2;255;105;180m", "coral": "\033[38;2;255;127;80m", "lavender": "\033[38;2;230;230;250m", "medium spring green": "\033[38;2;0;250;154m", "royal blue": "\033[38;2;65;105;225m",
		"chocolate": "\033[38;2;210;105;30m", "snow": "\033[38;2;255;250;250m", "slate gray": "\033[38;2;112;128;144m", "peru": "\033[38;2;205;133;63m", "turquoise": "\033[38;2;64;224;208m", "violet": "\033[38;2;238;130;238m",
		"khaki": "\033[38;2;240;230;140m", "sky blue": "\033[38;2;135;206;235m", "misty rose": "\033[38;2;255;228;225m",
	}

	return colors[color]
}

func getRgbColor(colorValue string) string {
	// getColor("451263")
	rgb := ""
	if len(colorValue) >= 4 && strings.ToLower(colorValue[:4]) == "rgb(" && string(colorValue[len(colorValue)-1:]) == ")" {
		rgb = formatRGBToAnsi(colorValue)
	} else if colorValue[:1] == "#" {
		r, _ := regexp.Compile(`[a-f0-9]+`)
		value := r.FindString(colorValue)
		if len(value) == len(colorValue[1:]) {
			///this if used to check if the user entersan invalid color zx : #3155q7, #54t451, #895x54
			rgb = formatRGBToAnsi(fromHex(value))
		}
	}

	return rgb
}

type AutoGenerated2 struct {
	Rgb struct {
		Value string `json:"value"`
	} `json:"rgb"`
}

// Get RGB color from Hex code system
func fromHex(color string) string {
	fmt.Println("color[1:] : ", color[1:])
	// myUrl := "https://csscolorsapi.com/api/colors/" + color
	myUrl := "https://www.thecolorapi.com/id?hex=" + color[0:]
	res, err := http.Get(myUrl)
	if err != nil {
		fmt.Println("err")
	}

	content, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err reading the response")
		os.Exit(1)
	}
	var myrgb AutoGenerated2
	json.Unmarshal(content, &myrgb)
	// fmt.Println(json.un)
	fmt.Println("res:", myrgb.Rgb.Value)
	return myrgb.Rgb.Value
}

func formatRGBToAnsi(colorValue string) string {
	///handle spaces
	r, _ := regexp.Compile(`[0-9]+\s*,\s*[0-9]+\s*,\s*[0-9]+`)
	if strings.Contains(colorValue, " ") {
		fmt.Println("contains")
		colorValue = strings.ReplaceAll(colorValue, " ", "")
	}
	// r, _ := regexp.Compile(`[0-9]+,[0-9]+,[0-9]+`)
	onlyNbr := r.FindString(colorValue)
	fmt.Println("onl", onlyNbr)
	if len(onlyNbr) >= 1 {
		///rgb ansi format
		return "\033[38;2;" + strings.ReplaceAll(onlyNbr, ",", ";") + "m"
	}
	return ""
}