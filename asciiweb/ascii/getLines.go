package ascii

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func GetLines(bannerFile string) []string {
	fmt.Println("bannerFile : ", bannerFile)

	///handle args error
	if bannerFile != "" {
		bannerFile = "./Banners/" + bannerFile + ".txt"
	} else {
		bannerFile = "./Banners/standard.txt"
	}

	file, err := os.ReadFile(bannerFile)

	if err != nil {
		fmt.Println("aError :", bannerFile, "file not found")

		return []string{}
	}
	lines := strings.Split(string(file), "\n")

	return lines
}

func RemoveSpaces(oldString string) string {
	regRule := regexp.MustCompile(`\s+`)
	return regRule.ReplaceAllString(oldString, " ")
}
