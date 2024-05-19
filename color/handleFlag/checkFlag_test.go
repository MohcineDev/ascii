package handleFlag

import (
	"fmt"
	"testing"
)

func TestIsValidFlagOutput(t *testing.T) {
	a, b, _, _ := IsValidFlag([]string{"--output=res.txt"})
	fmt.Println(a, b)
}

func TestIsValidFlagColor(t *testing.T) {
	_, _, c, d := IsValidFlag([]string{"--color=rgb(45,65,84)"})
	fmt.Println(c, d)
}

func TestIsValidTwoFlags(t *testing.T) {
	IsValidFlag([]string{"--output=res.txt", "--color=red"})
}
