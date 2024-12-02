package main

import (
	"fmt"
	"strings"

	"github.com/johnverrone/aoc2024/util"
)

const sample = `
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`

func main() {
	in := util.ParseInput("")
	lines := strings.Split(in, "\n")

	reports := [][]int{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		strLevels := strings.Fields(line)
		if len(strLevels) == 0 {
			continue
		}
		levels := []int{}
		for _, l := range strLevels {
			levels = append(levels, util.MustInt(l))
		}
		reports = append(reports, levels)
	}

	// part 1
	safeCount := 0
	for _, levels := range reports {
		if isSafe(levels) {
			safeCount++
		}
	}
	fmt.Println("Part 1:", safeCount)

	// part 2
	tolerantSafe := 0
	for _, levels := range reports {
		if isSafe(levels) {
			tolerantSafe++
		} else {
			for i := 0; i < len(levels); i++ {
				removed := removeIndex(levels, i)
				if isSafe(removed) {
					tolerantSafe++
					break
				}
			}
		}
	}
	fmt.Println("Part 2:", tolerantSafe)
}

func isSafe(r []int) bool {
	i := 1
	inc := r[0] < r[1]
	for {
		if i >= len(r) {
			break
		}
		c, p := r[i], r[i-1]
		if (c == p) ||
			(c < p && inc) ||
			(c > p && !inc) ||
			(util.Abs(c-p) > 3) {
			return false
		}
		i++
	}
	return true
}

func removeIndex(s []int, index int) []int {
	n := []int{}
	for i, v := range s {
		if i == index {
			continue
		}
		n = append(n, v)
	}
	return n
}
