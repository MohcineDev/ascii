package handleFlag

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

var (
	IsOutput = false
	IsColor  = false

	OutputFile = ""
	Color      = ""
)

// make sure only one flag is used
// return that flag and it's value
func IsValidFlag(myFlags []string) {
	checkIfOutput(myFlags[0])

	checkIfColor(myFlags[0])

}

func checkIfOutput(myFlag string) {
	if len(myFlag) > 13 && myFlag[:9] == "--output=" && path.Ext(myFlag[9:]) == ".txt" {
		OutputFile = myFlag[9:]
		IsOutput = true
	}
}

func checkIfColor(myFlag string) {

	if len(myFlag) >= 8 && myFlag[:8] == "--color=" {
		Color = myFlag[8:]

		IsColor = true
	}

	if IsColor {
		if strings.Contains(Color, "(") || strings.Contains(Color, "#") {
			checColorType(Color)
		} else {
			getANSIColor(strings.ToLower(Color))
		}
	}

}

func getANSIColor(color string) {
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

	Color = colors[color]
}

func checColorType(colorValue string) {
	// getColor("451263")
	rgb := ""
	if len(colorValue) >= 4 && strings.ToLower(colorValue[:4]) == "rgb(" && string(colorValue[len(colorValue)-1:]) == ")" {
		rgb = formatRGBToAnsi(colorValue)
	} else if colorValue[:1] == "#" {
		r, _ := regexp.Compile(`[a-fA-F0-9]+`)
		value := r.FindString(colorValue)
		if len(value) == len(colorValue[1:]) {
			///this if used to check if the user enters an invalid color zx : #3155q7, #54t451, #895x54

			if len(value) > 6 {
				value = value[:6]
			}
			rgb = formatRGBToAnsi(fromHex(value))
		}
	}

	Color = rgb
}

type AutoGenerated2 struct {
	Rgb struct {
		Value string `json:"value"`
	} `json:"rgb"`
}

// Get RGB color from Hex code system
func fromHex(color string) string {
	myUrl := "https://www.thecolorapi.com/id?hex=" + color[0:]
	res, err := http.Get(myUrl)
	if err != nil {
		fmt.Println("err handling the request")
	}

	content, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("err reading the response")
		os.Exit(0)
	}
	var myrgb AutoGenerated2
	json.Unmarshal(content, &myrgb)
	return myrgb.Rgb.Value
}

func formatRGBToAnsi(colorValue string) string {
	///handle spaces
	r, _ := regexp.Compile(`[0-9]+\s*,\s*[0-9]+\s*,\s*[0-9]+`)
	if strings.Contains(colorValue, " ") {
		colorValue = strings.ReplaceAll(colorValue, " ", "")
	}
	onlyNbr := r.FindString(colorValue)
	if len(onlyNbr) >= 1 {
		///format RGB to ANSI format
		return "\033[38;2;" + strings.ReplaceAll(onlyNbr, ",", ";") + "m"
	}
	return ""
}
