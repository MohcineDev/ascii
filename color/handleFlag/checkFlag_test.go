package handleFlag

import (
	"testing"
)

func TestIsValidFlagOutput(t *testing.T) {
	IsValidFlag([]string{"--output=res.txt"})
}

func TestIsValidFlagColor(t *testing.T) {
	IsValidFlag([]string{"--color=rgb(45,65,84)"})
}

func TestIsValidTwoFlags(t *testing.T) {
	IsValidFlag([]string{"--output=res.txt", "--color=red"})
}
