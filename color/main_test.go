package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// / TODO
// add func runTest

func runTest(args []string, t *testing.T, file string) {
	cmd := exec.Command("./main", args...)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("execution error")
	}
	content, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf(err.Error())
	}

	///if output doesn't equal the content of file / string
	if !strings.EqualFold(string(content), string(output)) {
		t.Fatalf("not equal")
	}
}

func TestMainOneArg(t *testing.T) {
	args := []string{"1"}
	runTest(args, t, "./test/one.txt")
}

// /wrong banner file
func TestMainTwoArgWrongFile(t *testing.T) {
	args := []string{"1", "o"}
	runTest(args, t, "./test/wrongFile.txt")
}

// /correct banner file
func TestMainTwoArgsCorrectFile(t *testing.T) {
	args := []string{"A", "shadow"}
	runTest(args, t, "./test/CorrectFile.txt")
}

// /compare two files --output FLAG
func runTestTwoFile(args []string, t *testing.T, srcfile string, resfile string) {
	cmd := exec.Command("./main", args...)
	cmd.Output()

	src, srcErr := os.ReadFile(srcfile)
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
}

func TestMainTwoArgsCorrectFlag(t *testing.T) {
	args := []string{"--output=res.txt", "shadow"}
	runTestTwoFile(args, t, "./test/src.txt", "./res.txt")
}

func TestMainTwoArgsIncorrectFlag(t *testing.T) {
	cmd := exec.Command("./main", "-output=res.txt", "thinkertoy")

	output, _ := cmd.Output()

	src, err0 := os.ReadFile("./test/IncorrectFlag.txt")
	if err0 != nil {
		t.Fatalf("error")
	}

	if !strings.EqualFold(string(output), string(src)) {
		t.Fatalf("not equal")
	}
}

/////THREE ARGS

func TestMainThreeArgs(t *testing.T) {
	cmd := exec.Command("./main", "--output=./test/threeArgs.txt", "Hello", "thinkertoy")
	cmd.Output()

	res, err := os.ReadFile("./test/threeArgs.txt")
	if err != nil {
		t.Fatalf("error")
	}
	result, err := os.ReadFile("./test/threeArgsResult.txt")
	if err != nil {
		t.Fatalf("error")
	}
	if !strings.EqualFold(string(result), string(res)) {
		t.Fatalf("not equal")
	}
}

// more than 3 arguments
func TestMainMoreThanThreeArgs(t *testing.T) {
	cmd := exec.Command("./main", "--output=./test/four.txt", "Hello", "thinkertoy", "sz")

	output, _ := cmd.Output()

	usageMsg := "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n"

	if !strings.EqualFold(usageMsg, string(output)) {
		t.Fatalf("error")
	}
}
