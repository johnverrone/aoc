package util

import "testing"

func TestParseInput(t *testing.T) {
	in := ParseInput()

	if in == "" {
		t.Errorf("Error parsing input from input.txt.")
	}
}
