package main

import (
	"fmt"
	"strings"

	"github.com/johnverrone/aoc2025/util"
)

const sample = `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func main() {
	in := util.ParseInput(sample)
	lines := strings.SplitSeq(in, "\n")

	for line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Println(line)
	}
}
