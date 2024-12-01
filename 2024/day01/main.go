package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/johnverrone/aoc2024/util"
)

const sample = `
3   4
4   3
2   5
1   3
3   9
3   3
`

func main() {
	in := util.ParseInput(nil)
	lines := strings.Split(in, "\n")

	left := []int{}
	right := []int{}
	rightCount := map[int]int{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) != 2 {
			panic("could not parse input")
		}
		left = append(left, mustInt(parts[0]))
		rv := mustInt(parts[1])
		right = append(right, rv)
		rightCount[rv] = rightCount[rv] + 1
	}

	// part 1
	// sort the lists
	sort.Sort(sort.IntSlice(left))
	sort.Sort(sort.IntSlice(right))

	dist := 0
	for i := 0; i < len(left); i++ {
		dist += abs(right[i] - left[i])
	}

	fmt.Println("Part 1:", dist)

	// part 2
	ss := 0
	for _, v := range left {
		//lookup the count in right
		rc := rightCount[v]
		ss += rc * v
	}

	fmt.Println("Part 2:", ss)
}

func mustInt(s string) int {
	if v, err := strconv.Atoi(s); err != nil {
		panic(fmt.Sprint(s, err))
	} else {
		return v
	}
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
