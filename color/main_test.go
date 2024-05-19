package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

var (
	outputUsageMessage = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n"
	colorUsageMessage  = "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"\n"
)

func compareTwoStrings(args []string, t *testing.T, usageMsg string) {
	cmd := exec.Command("./main", args...)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("execution error")
	}
	///if output doesn't equal the content of file / string
	if !strings.EqualFold(usageMsg, string(output)) {
		t.Fatalf("not equal")
	}
	fmt.Println("")
}

func CompareFileWithString(args []string, t *testing.T, file string) {
	cmd := exec.Command("./main", args...)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("execution error")
	}
	content, err := os.ReadFile("./test/" + file)
	if err != nil {
		t.Fatalf(err.Error())
	}
	///if output doesn't equal the content of file / string
	if !strings.EqualFold(string(content), string(output)) {
		t.Fatalf("not equal")
	}
	fmt.Println("")
}

// /compare tcolorwo files --output FLAG
func compareTwoFiles(args []string, t *testing.T, srcfile string, resfile string) {
	cmd := exec.Command("./main", args...)
	cmd.Output()

	src, srcErr := os.ReadFile("./test/" + srcfile)
	if srcErr != nil {
		t.Fatalf(srcErr.Error())
	}
	res, resErr := os.ReadFile(resfile)
	if resErr != nil {
		t.Fatalf(resErr.Error())
	}

	///if output doesn't equal the content of file / string
	if !strings.EqualFold(string(src), string(res)) {
		t.Fatalf("not equal")
	}
	fmt.Println("")
}

// ///////////
// ///////////
func TestMainOneArg(t *testing.T) {
	args := []string{"1"}
	CompareFileWithString(args, t, "one.txt")
}

// /wrong banner file
func TestMainTwoArgWrongFile(t *testing.T) {
	args := []string{"1", "ogreg"}
	compareTwoStrings(args, t, "Error : ../ogreg.txt file not found\n")

	// CompareFileWithString(args, t, "wrongFile.txt")
}

// /correct banner file
func TestMainTwoArgsCorrectFile(t *testing.T) {
	args := []string{"A", "shadow"}
	CompareFileWithString(args, t, "CorrectFile.txt")
}

// //

func TestMainTwoArgsCorrectFlag(t *testing.T) {
	args := []string{"--output=res.txt", "shadow"}
	compareTwoFiles(args, t, "src.txt", "res.txt")
}

func TestMainTwoArgsIncorrectFlag(t *testing.T) {
	args := []string{"-output=res.txt", "thinkertoy"}
	compareTwoStrings(args, t, outputUsageMessage)
}

/////THREE ARGS
// /TEST OUTPUT FLAG

func TestMainThreeArgs(t *testing.T) {
	args := []string{"--output=./res/threeArgsResult.txt", "Hello", "thinkertoy"}
	compareTwoFiles(args, t, "threeArgs.txt", "threeArgsResult.txt")
}

// more than 3 arguments
func TestMainMoreThanThreeArgs(t *testing.T) {
	args := []string{"--output=four.txt", "Hello", "thinkertoy", "sz"}
	// 1 - args - 3 default display colorUsageMessage
	compareTwoStrings(args, t, colorUsageMessage)
}

// /TEST COLOR FLAG
func TestMainValidFlag(t *testing.T) {
	args := []string{"--color=red", "1"}
	CompareFileWithString(args, t, "one.txt")
}

func TestMainInvalidFlag(t *testing.T) {
	args := []string{"-color=red", "hello"}
	compareTwoStrings(args, t, colorUsageMessage)
}

func TestMainInvalidFlag0(t *testing.T) {
	args := []string{"--clor=red", "hello"}
	compareTwoStrings(args, t, colorUsageMessage)
}

func TestMainInvalidFlag1(t *testing.T) {
	args := []string{"--color+red", "hello"}
	compareTwoStrings(args, t, colorUsageMessage)
}
