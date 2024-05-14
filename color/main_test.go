package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestMainOneArg(t *testing.T) {
	cmd := exec.Command("./output", "1")
	content, err := os.ReadFile("./test/one.txt")
	if err != nil {
		t.Fatalf("error")
	}
	output, _ := cmd.Output()
	if !strings.EqualFold(string(content), string(output)) {
		t.Fatalf("not equal")
	}
}

func TestMainTwoArgWrongFile(t *testing.T) {
	cmd := exec.Command("./output", "1", "a")

	output, _ := cmd.Output()
	if !strings.EqualFold("Error : a.txt file not found\n", string(output)) {
		t.Fatalf("not equal")
	}
}

func TestMainTwoArgsCorrectFile(t *testing.T) {
	cmd := exec.Command("./output", "A", "shadow")

	output, _ := cmd.Output()

	content, err := os.ReadFile("./test/CorrectFile.txt")
	if err != nil {
		t.Fatalf("error")
	}

	if !strings.EqualFold(string(content), string(output)) {
		t.Fatalf("not equal")
	}
}

func TestMainTwoArgsCorrectFlag(t *testing.T) {
	cmd := exec.Command("./output", "--output=res.txt", "shadow")

	cmd.Output()

	res, err := os.ReadFile("res.txt")
	if err != nil {
		t.Fatalf("error")
	}
	src, err0 := os.ReadFile("./test/CorrectFlag.txt")
	if err0 != nil {
		t.Fatalf("error")
	}

	if !strings.EqualFold(string(res), string(src)) {
		t.Fatalf("not equal")
	}
}

func TestMainTwoArgsIncorrectFlag(t *testing.T) {
	cmd := exec.Command("./output", "T", "thinkertoy")

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
	cmd := exec.Command("./output", "--output=./test/threeArgs.txt", "Hello", "thinkertoy")
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
	cmd := exec.Command("./output", "--output=./test/four.txt", "Hello", "thinkertoy", "sz")

	output, _ := cmd.Output()

	usageMsg := "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n"

	if !strings.EqualFold(usageMsg, string(output)) {
		t.Fatalf("error")
	}
}
